package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("mysql", "root:Sa123456@tcp(localhost:3306)/sesame")
	//db, err := sql.Open("sqlite3", "./sesame.db")
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")

	if err != nil {
		fmt.Println("Open Fail", err)
		return
	}

	defer db.Close()
	//stmt, err := db.Prepare("SELECT * FROM USER WHERE userId=?")
	//注意: pgsql参数是用$1, $2这类代替, 和mysql, sqlite不同
	stmt, err := db.Prepare("SELECT userid, name, age FROM MyUser WHERE userid=$1")

	if err != nil {
		fmt.Println("Prepare Fail ", err)
		return
	}

	//查询
	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Println("Query Fail", err)
		return
	}

	for rows.Next() {
		var uid int
		var username string
		var age int

		err = rows.Scan(&uid, &username, &age)
		if err == nil {
			fmt.Println(uid)
			fmt.Println(username)
			fmt.Println(age)
		}
	}

}
