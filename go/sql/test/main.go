package main

import (
	"os"
	"log"
	"strconv"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	orm "github.com/mojosa-software/gorm"
	"github.com/mojosa-software/gorm/driver/mysql"
)

type Todo struct {
	orm.Model
	Label string
	Text string
}

func main() {
	portStr := os.Getenv("MYSQL_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}

	connConfig := orm.ConnConfig{
		Driver: "mysql",
		Protocol: "tcp",
		Login:    os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     port,
		Name:     os.Getenv("MYSQL_DB"),
		ParseTime: true,
	}
	dsn := connConfig.DsnString()
	db, err := orm.Open(mysql.Open(dsn), &orm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Todo{})
	db.Create(&Todo{
		Label: "The first todo!",
		Text: "Finish migrating the GORM",
	})
	todo := Todo{}
	db.First(&todo)
	log.Println("The todo:", todo)
}
