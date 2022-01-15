package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"ms-address/models"
)

//type Database struct {
//	DB *gorm.DB
//}

func Init(databaseURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	e := db.AutoMigrate(&models.Address{}, &models.Transactions{})
	if e != nil {
		return nil
	}
	return db
}
