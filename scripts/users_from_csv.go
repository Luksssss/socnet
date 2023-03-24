package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload" // load env vars from .env file
	"github.com/kelseyhightower/envconfig"
	"otus/socNet/internal/hash"
	"otus/socNet/internal/structs"
)

// ConfigScript represents application configuration.
type ConfigScript struct {
	PGHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PGPort     uint16 `envconfig:"POSTGRES_PORT" required:"true"`
	PGDatabase string `envconfig:"POSTGRES_DATABASE" required:"true"`
	PGParams   string `envconfig:"POSTGRES_PARAMS"`
	PGUsername string `envconfig:"POSTGRES_USERNAME" required:"true"`
	PGPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
}

func main() {
	fmt.Println("start.")
	ctx := context.Background()
	cfg, err := readConfigScript()
	if err != nil {
		panic(fmt.Sprintf("configuration error: %v", err))
	}
	pool, err := getDBClient(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// пишем всё в одной транзакции (для ускорения)
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})

	url := "https://raw.githubusercontent.com/OtusTeam/highload/master/homework/people.csv"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	clearUsers(ctx, pool)

	count := 0
	for _, user := range data {
		if len(user) > 0 {
			names := strings.Split(user[0], " ")
			if len(names) < 2 {
				continue
			}
			userData := &structs.User{
				FirstName:  names[1],
				SecondName: names[0],
				DateBirth:  convDateBirth(user[1]),
				City:       user[2],
				//Pass:       generatePass(),
				Biography: generateBio(),
			}
			saveUser(ctx, tx, userData)
		}
		count++
		if count%50000 == 0 {
			fmt.Printf("add rows = %d\n", count)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish.")
}

func generateBio() string {
	in := []string{"футбол, настольные игры", "чтение, прогулки", "сериалы, сквош", "фильмы"}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	return pick
}

func convDateBirth(years string) time.Time {
	y, err := strconv.Atoi(years)
	if err != nil {
		y = 20
	}
	date := time.Now().AddDate(-y, 0, 0)
	return date
}

// оч сильно замедляет вставку (!)
func generatePass() string {
	pass, _ := hash.HashPassword("otus")
	return pass
}

func saveUser(ctx context.Context, tx pgx.Tx, userData *structs.User) {
	var (
		idUser int64
	)
	query := `insert into users (first_name, second_name, date_birth, id_city, pass)
			  values (
			          $1,
			          $2, 
			          $3, 
			          (select id from city where name = $4),
			          $5
			          ) returning id`

	err := tx.QueryRow(ctx, query,
		userData.FirstName,
		userData.SecondName,
		userData.DateBirth,
		userData.City,
		userData.Pass,
	).Scan(&idUser)

	if err != nil {
		panic(err)
	}
}

func clearUsers(ctx context.Context, pool *pgxpool.Pool) {
	query := `TRUNCATE TABLE users`

	_, err := pool.Exec(ctx, query)

	if err != nil {
		panic(err)
	}
	fmt.Println("clear table.")
}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getDBClient(ctx context.Context, cfg *ConfigScript) (*pgxpool.Pool, error) {
	DBUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PGUsername, cfg.PGPassword,
		cfg.PGHost, cfg.PGPort,
		cfg.PGDatabase,
	)
	conn, err := pgxpool.ParseConfig(DBUrl)
	conn.MaxConns = 100
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(ctx, conn)
	if err != nil {
		return nil, err
	}
	if err = checkConnectivity(ctx, pool); err != nil {
		return nil, err
	}

	return pool, nil
}

func checkConnectivity(ctx context.Context, pool *pgxpool.Pool) error {
	var err error
	for i := 0; i < 3; i++ {
		cctx, cancel := context.WithTimeout(ctx, time.Second)
		err = pool.Ping(cctx)
		if err != nil {
			time.Sleep(time.Second)
			cancel()
		} else {
			cancel()
			return nil
		}
	}
	return err
}

func readConfigScript() (*ConfigScript, error) {
	cfg := &ConfigScript{}
	err := envconfig.Process("", cfg)
	return cfg, err
}
