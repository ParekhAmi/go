package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Database Connection")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called db_1
	con, err := sql.Open("mysql", "root:example@/db_1")

	if con != nil {
		fmt.Println("Connection is done...")
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer con.Close()
}
