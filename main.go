package main

import (
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "randipsql"
	dbPass := "RandipBhachu94!"
	dbName := "bhachuDB"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func readAll(db *sql.DB) {
	
	// query
	sel, err := db.Query("SELECT * FROM people")
	if err != nil {
		panic(err.Error())
	}

	// query table
	for sel.Next() {
		var person string
		var weight int
		var date string

		err = sel.Scan(&person, &weight, &date)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("person: %s, weight: %d, date: %s\n", person, weight, date)
	}
}

func insert(person string, weight int, date string, db *sql.DB) {
	// TODO: insert in table

	// INSERT INTO DB
	// prepare  
	stmt, err := db.Prepare("INSERT INTO people(person, weight, date) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
 
	//execute
	res, err := stmt.Exec(person, weight, date)
	if err != nil {
		panic(err.Error())
	}	
 
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
 
	fmt.Println("Insert id", id)

	fmt.Println("Data to be added in database:")
	fmt.Printf("person: %s, weight: %d, date: %s\n", person, weight, date)
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db := dbConn()

	fmt.Println("Welcome to weight tracker:\n")
	for {
		fmt.Print("Enter [1] to record a new entry, [2] to read all entries [q] to  quit: ")
		var line, name, weight, date string
		fmt.Scanln(&line)
		// 
		if line == "q" {
			break
		} else if line == "1" {
			fmt.Print("name: ")
			fmt.Scanln(&name)
			fmt.Print("weight: ")
			fmt.Scanln(&weight)
			fmt.Print("date (yyyy-mm-dd): ")
			fmt.Scanln(&date)
			w, _ := strconv.ParseInt(weight, 0, 32)
			insert(name, int(w), date, db);
		} else if line == "2" {
			readAll(db)
		}

	}


	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
