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

type Config struct {
	DbSource      string `env:"DB_SOURCE"`
	DbType        string `env:"DB_TYPE"`
	DbUsername    string `env:"DB_USERNAME"`
	DbPassword    string `env:"DB_PASSWORD"`
	DbName        string `env:"DB_NAME"`
	DbTestName    string `env:"DB_TEST_NAME"`
	DbHost        string `env:"DB_HOST" env-default:"localhost"`
	DbPort        string `env:"DB_PORT" env-default:"5432"`
	DbSSLMode     string `env:"DB_SSL_MODE" env-default:"false"`
	ServerAddress string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8888"`
}

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

	// dbConn, err := sql.Open(config.DbType, dbSource)

	poolConfig, err := pgxpool.ParseConfig(dbSource)

	poolConfig.MaxConnLifetime = 100

	dbConn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

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
