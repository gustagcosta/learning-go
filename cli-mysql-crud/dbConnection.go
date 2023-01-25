package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:docker@/cli-mysql-crud")
	if err != nil {
		panic(err.Error())
	}
	return db
}
