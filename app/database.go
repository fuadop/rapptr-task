package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func ConnectDatabase() *sql.DB {
	dsn := getDsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db
}

func getDBCredentials() (string, string, string, string, string) {
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		dbUser = "root"
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = "3306"
	}

	return dbUser, dbPass, dbHost, dbPort, dbName
}

func formatDsn(user, pass, host, port, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pass, host, port, dbName)
}

func getDsn() string {
	dbUser, dbPass, dbHost, dbPort, dbName := getDBCredentials()
	return formatDsn(dbUser, dbPass, dbHost, dbPort, dbName)
}

func (a *App) AutoMigrate() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(250) NOT NULL UNIQUE,	
			username VARCHAR(250) NOT NULL UNIQUE,
			password VARCHAR(250) NOT NULL,
			createdAt DATETIME NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id)
		);
		CREATE TABLE IF NOT EXISTS todos (
			id VARCHAR(250) NOT NULL UNIQUE,	
			userId VARCHAR(250) NOT NULL,
			text VARCHAR(250) NOT NULL,
			isDeleted BOOL NOT NULL DEFAULT FALSE,
			createdAt DATETIME NOT NULL DEFAULT NOW(),
			updatedAt DATETIME NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id)
		);
	`

	for _, ddl := range strings.Split(query, ";") {
		if len(ddl) <= 3 {
			continue
		}

		if _, err := a.DB.Exec(ddl); err != nil {
			log.Fatalln(err)
		}
	}
}
