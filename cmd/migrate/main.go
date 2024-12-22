package main

import (
	"database/sql"
	"log"

	"github.com/go-rest-api/cmd/api"
	"github.com/go-rest-api/config"
	"github.com/go-rest-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		// Will delete later
		//Can't put this too production hehe!
		User:                 config.Environs.DBUser,
		Passwd:               config.Environs.DBPassword,
		Addr:                 config.Environs.DBAddress,
		DBName:               config.Environs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Bros! The db works!!!")
}
