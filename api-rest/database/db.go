package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	DB, err = gorm.Open(mysql.Open("root:docker@/api-rest-go"))
	if err != nil {
		log.Panic("error connecting to database")
	}
}
