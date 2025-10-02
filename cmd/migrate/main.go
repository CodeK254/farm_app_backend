package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CodeK254/farm_app_backend/config"
	"github.com/CodeK254/farm_app_backend/db"
	mySQLConfig "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewSQLStorage(mySQLConfig.Config{
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "force" {
		if len(os.Args) < 3 {
			log.Fatal("force command requires a version number")
		}
		version := os.Args[len(os.Args)-2]

		// Parse version string to int
		var versionInt int
		if _, err := fmt.Sscanf(version, "%d", &versionInt); err != nil {
			log.Fatal("invalid version number:", err)
		}

		if err := m.Force(versionInt); err != nil {
			log.Fatal(err)
		}
	}
}
