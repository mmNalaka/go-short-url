package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/mmnalaka/go-short-url/pkg/utils"
)

const (
	dbUserName = "postgres"
	dbPassword = "postgres"
	dbName     = "short_url"
	dbHost     = "localhost"
	dbPort     = "5432"
)

var (
	Client   *sql.DB
	username = utils.GetEnv("POSTGRES_USERNAME", dbUserName)
	password = utils.GetEnv("POSTGRES_PASSWORD", dbPassword)
	host     = utils.GetEnv("POSTGRES_HOST", dbHost)
	port     = utils.GetEnv("POSTGRES_PORT", dbPort)
	name     = utils.GetEnv("POSTGRES_NAME", dbName)
)

func Init() {
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, name)

	var err error
	Client, err = sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to database")
}

func RunMigrations() {
	migration := `
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			long_url TEXT NOT NULL,
			short_url TEXT
		);
	`
	_, err := Client.Exec(migration)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully ran migrations")
}
