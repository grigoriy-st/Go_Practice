package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := `
				user=postgres
				dbname=my_db
				password=0000
				host=localhost
				sslmode=disable
				`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
					CREATE TABLE IF NOT EXISTS users(
						id SERIAL PRIMARY KEY,
						name VARCHAR(100),
						email VARCHAR(100) UNIQUE NOT NULL
					);
						`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Table is create succesfull")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Insert your email: ")
	scanner.Scan()
	email := scanner.Text()

	insertSQL := `
				INSERT INTO users
					(name, email)
				VALUES
					($1, $2) RETURNING id;
				`
	var userID int
	err = db.QueryRow(insertSQL, name, email).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("New user %s added", name)
	}
	// print(name, email)
}
