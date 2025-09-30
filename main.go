package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CodeK254/farm_app_backend/cmd/api"
	"github.com/CodeK254/farm_app_backend/config"
	"github.com/CodeK254/farm_app_backend/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil{
		log.Fatalf("Error setting up DB %v", err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal("error configuring server ", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil{
		log.Fatalf("failed to connect to DB %v", err)
	}

	fmt.Println("connection to database successfully")
}