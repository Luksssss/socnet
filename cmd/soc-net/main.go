package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"otus/socNet/database"
	database_adapter "otus/socNet/internal/adapter/database"
	socnet_adapter "otus/socNet/internal/adapter/soc_net"
	"otus/socNet/internal/app/soc_net"
	pb "otus/socNet/internal/pb/api/socnet"
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

	go runRest(ctx, cfg)

	s := grpc.NewServer()
	pb.RegisterSocNetServer(s, i)
	log.Printf("Starting gRPC listener on port %d", cfg.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initService(ctx context.Context, cfg *Config) (soc_net.SocNetAPI, error) {
	DBUrlMaster := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PGUsername, cfg.PGPassword,
		cfg.PGHost, cfg.PGPort,
		cfg.PGDatabase,
	)

	connMaster, err := pgxpool.ParseConfig(DBUrlMaster)
	if err != nil {
		return nil, err
	}

	DBUrlSlave := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PGUsername, cfg.PGPassword,
		cfg.PGHostSlave, cfg.PGPort,
		cfg.PGDatabase,
	)

	connSlave, err := pgxpool.ParseConfig(DBUrlSlave)
	if err != nil {
		return nil, err
	}
	//conn.MaxConns = 6
	//conn.ConnConfig.TLSConfig = nil
	poolMaster, err := pgxpool.ConnectConfig(ctx, connMaster)
	if err != nil {
		return nil, err
	}
	if err = checkConnectivity(ctx, poolMaster); err != nil {
		return nil, err
	}
	poolSlave, err := pgxpool.ConnectConfig(ctx, connSlave)
	if err != nil {
		return nil, err
	}
	if err = checkConnectivity(ctx, poolSlave); err != nil {
		return nil, err
	}
	dbClient := database.New(poolMaster, poolSlave)

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

func runRest(ctx context.Context, cfg *Config) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterSocNetHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", cfg.Port), opts)
	if err != nil {
		panic(err)
	}
	log.Printf(fmt.Sprintf("server http listening at %d", cfg.HttpPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), mux); err != nil {
		panic(err)
	}
}
