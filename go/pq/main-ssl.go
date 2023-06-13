package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host         = "ep-billowing-sun-767748.us-west-2.aws.neon.tech"
	port         = 5432
	user         = "<user>"
	password     = "<password>"
	dbname       = "<dbname>"
	sslrootcert  = "/path/to/root.crt"
)

func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full&sslrootcert=%s", user, password, host, dbname, sslrootcert)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	var version string
	db.QueryRow("SELECT version()").Scan(&version)
	fmt.Println("PostgreSQL version:", version)

	var currentTime string
	db.QueryRow("SELECT CURRENT_TIMESTAMP").Scan(&currentTime)
	fmt.Println("Current Timestamp:", currentTime)
}
