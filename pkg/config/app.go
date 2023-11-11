package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

   var db *gorm.DB
   

   func Connect () {
	dbURL := "postgres://postgres:postgres@localhost:5432/go-postgres"
	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = d
   }


   func GetDB() *gorm.DB {
	return db
   }