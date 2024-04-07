package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/shammianand/ecom/cmd/api"
	"github.com/shammianand/ecom/config"
	"github.com/shammianand/ecom/db"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	addr := fmt.Sprintf(":%s", config.Envs.Port)
	server := api.NewAPIServer(addr, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to DB")

}
