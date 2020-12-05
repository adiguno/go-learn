package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test-go-db"
)

func main() {
	// ssl is disabled because lib/pq does not have it enabled by default
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// .Open() validates the connection info
	// (driver name, connection info for driver to connect)
	// returns a pointer to db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// defers the db.Close() until surrounding function returns
	// here: close db connection after everything
	defer db.Close()

	// .Ping() actually creates a connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
