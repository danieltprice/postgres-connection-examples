package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

const (
	host         = "<host>"
	port         = 5432
	user         = "<user>"
	password     = "<password>"
	dbname       = "<dbname>"
)

func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", user, password, host, dbname)

	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to connect to database: %v\n", err))
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected!")

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PostgreSQL version:", version)

	var currentTime string
	err = conn.QueryRow(context.Background(), "SELECT CURRENT_TIMESTAMP::text").Scan(&currentTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Timestamp:", currentTime)
}

