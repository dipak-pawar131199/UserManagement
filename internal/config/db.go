package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

type DbConnection struct {
	DB *sql.DB
}

var dbInstance *DbConnection

func ConnectionDB() *DbConnection {

	once.Do(func() {

		db, err := sql.Open("mysql", "root:redhat@tcp(127.0.0.1:3306)/user")

		if err != nil {
			fmt.Println("DB Open Failed:", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal("Failed to ping database")
		}
		dbInstance = &DbConnection{DB: db}
	})
	return dbInstance

}
