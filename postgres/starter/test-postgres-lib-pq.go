package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test-go-db"
)

type User struct {
	username string
	uuid     string
	points   int
}

// func insert_user(db DB, user User) {
// 	query := fmt.Sprintf("insert into all_user(\"username\", \"uuid\", \"points\") values (%s, %s, %d)", user.username, user.uuid, user.points)

// 	db.Exec()
// }

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

	goUser := User{
		username: "go-asdf",
		uuid:     uuid.New().String(),
		points:   10,
	}
	fmt.Println(goUser.username + " - " + goUser.uuid + " - " + strconv.Itoa(goUser.points))

	// insert_user(user)
	query := fmt.Sprintf("insert into all_user(\"USERNAME\", \"UUID\", \"POINTS\") values ('%s', '%s', %d)", goUser.username, goUser.uuid, goUser.points)
	fmt.Println(query)

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}
