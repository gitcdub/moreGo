// this is an example of using go with a postgres database
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
	password = "!QW#eft6yhu8"
	dbname   = "test_group"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO emp (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "jim@fib.com", "jim", "smith").Scan(&id)
	if err != nil {
		panic(err)
	}

	/* This is for the inital connection test to the database.
	   	err = db.Ping()
	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Println("Successfully connected!")
	   }
	*/

	fmt.Println("New record ID is:", id)
}
