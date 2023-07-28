package main

import (
	"os"
	"fmt"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	addr := fmt.Sprintf("127.0.0.1:%s", os.Getenv("MYSQL_PORT"))
	fmt.Printf("%q", addr)
	cfg := mysql.Config{
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net: "tcp",
		Addr: addr,
		DBName: os.Getenv("MYSQL_DB"),
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}

