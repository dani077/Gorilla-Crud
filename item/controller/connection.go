package controller

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Con struct {
	DB *sql.DB
}

func DBConnection(*sql.DB) *sql.DB {
	db, err := sql.Open("mysql", "root:root@/dbselling")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
