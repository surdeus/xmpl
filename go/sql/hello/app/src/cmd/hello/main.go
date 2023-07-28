package main

import (
	"os"
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mojosa-software/gosrv/src/dbx/sqlx"
	"strconv"
)

var db *sql.DB


func main() {
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	connString := sqlx.ConnConfig{
			Driver: "mysql",
			Login: os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host: os.Getenv("MYSQL_HOST"),
			Port: port, 
			Name: os.Getenv("MYSQL_DB"),
	}.String()
	fmt.Printf("%q\n", connString)
	
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Query("create table todos (Id int)")
	if err != nil {
		panic(fmt.Sprintf("creating table: %s", err.Error()))
	}
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}

