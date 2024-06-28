package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	PostgresHost     = "POSTGRESHOST"
	PostgresPort     = 1234 //POSTGRESPORT
	PostgresDb       = "POSTGRESDBNAME"
	PostgresUser     = "POSTGRESUSER"
	PostgresPassword = "POSTGRESPASSWORD"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresDb, PostgresUser, PostgresPassword)
	DB, err = sql.Open("postgres", config)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	log.Println("Connected to Database")
}

// func AutoMigrate() {
// 	ConnectDB()

// 	_, err := DB.Exec(`
// 	CREATE TABLE IF NOT EXISTS products (
// 		barcode INT PRIMARY KEY,
// 		name VARCHAR(255) NOT NULL,
// 		price INT NOT NULL,
// 		active BOOLEAN NOT NULL
//  	)
// 	`)

// 	if err != nil {
// 		log.Fatal("failed to create table:", err)
// 	}

// 	log.Println("table migrated successfully")
// }
