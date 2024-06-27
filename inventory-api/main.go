package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bobbybof/inventory-api/config"
	"github.com/bobbybof/inventory-api/internal/api"
	"github.com/bobbybof/inventory-api/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig(".env")

	if err != nil {
		log.Fatal("cannot load env: ", err)
	}

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
		config.DbSSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dbSource)

	if err != nil {
		log.Fatal("Faild parse config", err)
	}

	poolConfig.MaxConnLifetime = 100

	dbConn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		log.Fatal("Failed to open connection to database", err)
	}

	err = dbConn.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to ping database ", err)
	}

	defer func() {
		dbConn.Close()
	}()

	store := repository.NewStore(dbConn)

	server, err := api.NewServer(*config, store)

	if err != nil {
		log.Fatal("cannot make server")
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
