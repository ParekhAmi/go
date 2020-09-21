package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*type user struct {
	Name string
}*/

func main() {
	fmt.Println("Go MySQL Database Connection")

	// Open up our database connection.
	// The database is called db_1
	con, err := sql.Open("mysql", "root:example@tcp(172.19.0.2:3306)/db_1")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection established..")
	}

	defer con.Close()

	insert, err := con.Prepare("INSERT INTO Users(Id, Name) VALUES (1, 'test')")

	if err != nil {

		fmt.Println(err)
		//panic(err.Error())
	} else {
		fmt.Println("New record inserted")
	}

	defer insert.Close()

	_, err = insert.Exec(nil, "John")
	if err != nil {
		panic(err)
	}

	/*results, err := con.Query("SELECT UserName FROM Users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var User user

		err = results.Scan(&User.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(User.Name)

	}
	defer results.Close()*/
}
