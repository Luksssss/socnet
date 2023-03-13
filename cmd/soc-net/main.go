package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"otus/socNet/database"
	database_adapter "otus/socNet/internal/adapter/database"
	socnet_adapter "otus/socNet/internal/adapter/soc_net"
	"otus/socNet/internal/app/soc_net"
	pb "otus/socNet/internal/pb"
)

const pingAttemptCount = 3

func main() {
	ctx := context.Background()
	cfg, err := readConfig()
	if err != nil {
		panic(fmt.Sprintf("configuration error: %v", err))
	}
	log := initLogger(cfg.Development)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	adapter, err := initService(ctx, cfg)
	if err != nil {
		log.Fatalf("failed initService: %v", err)
	}
	i := soc_net.NewSocNetAPI(adapter)

	s := grpc.NewServer()
	pb.RegisterSocNetServer(s, i)
	log.Printf("Starting gRPC listener on port %d", cfg.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initService(ctx context.Context, cfg *Config) (soc_net.SocNetAPI, error) {
	DBUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PGUsername, cfg.PGPassword,
		cfg.PGHost, cfg.PGPort,
		cfg.PGDatabase,
	)

	conn, err := pgxpool.ParseConfig(DBUrl)
	if err != nil {
		return nil, err
	}
	//conn.MaxConns = 6
	//conn.ConnConfig.TLSConfig = nil
	pool, err := pgxpool.ConnectConfig(ctx, conn)
	if err != nil {
		return nil, err
	}
	if err = checkConnectivity(ctx, pool); err != nil {
		return nil, err
	}
	dbClient := database.New(pool)

	return socnet_adapter.New(
		database_adapter.New(dbClient),
	), nil
}

func initLogger(debug bool) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	if debug {
		log.Level = logrus.DebugLevel
		log.Warn("Debug is enabled")
	} else {
		log.Level = logrus.InfoLevel
	}
	return log
}

func checkConnectivity(ctx context.Context, pool *pgxpool.Pool) error {
	var err error
	for i := 0; i < pingAttemptCount; i++ {
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
