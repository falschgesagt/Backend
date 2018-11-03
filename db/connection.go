package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/falschgesagt/Backend/types"
	"github.com/falschgesagt/Backend/utilities"
	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func init() {

	config := types.DatabaseConfig{
		Host:     utilities.GetStringWithFallbackValueFrom("DB_HOST", "localhost"),
		Port:     utilities.GetIntWithFallbackValueFrom("DB_PORT", 3306),
		User:     utilities.GetStringFrom("DB_USER"),
		Password: utilities.GetStringFrom("DB_PASSWORD"),
		DB:       utilities.GetStringFrom("DB_NAME"),
	}

	// Builds the string for connection to database
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DB)
	connection, err := sql.Open("mysql", addr)

	if err != nil {
		log.Fatal(err)
	}

	utilities.ExecuteAndPanicOnError("CREATE TABLE IF NOT EXISTS quote (quote text)", connection)
	utilities.ExecuteAndPanicOnError("CREATE TABLE IF NOT EXISTS author (author text)", connection)

	Connection = connection
	log.Print("[db/connection.go] - Successfully connected!")
}