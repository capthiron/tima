package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {

	var dbDsn string
	switch os.Getenv("PROFILE") {
	case "dev":
		dbDsn = "./.tima.db"
	case "test":
		dbDsn = "file::memory:?cache=shared"
		fmt.Println("test database initialized")
	default:
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dbDsn = homedir + "/.tima.db"
	}

	db, err := gorm.Open(sqlite.Open(dbDsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect database:\n%v", err.Error())
		os.Exit(1)
	}

	if os.Getenv("PROFILE") == "dev" {
		db.Debug()
	}

	return db
}
