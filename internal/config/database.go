package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dns := "host=localhost user=postgres password=ows1234 dbname=salman_go port=5432 sslmode=disable sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	fmt.Println("Connected to Database")
	return db
}
