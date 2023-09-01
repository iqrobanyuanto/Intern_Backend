package config

import (
	"Intern_Backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	/*
		username := "root"
		password := "root"        // change this to your local db
		host := "localhost: 3306" // change this to your local db
		database := "dbintern"    // change this to your local db
	*/

	db, err := gorm.Open(mysql.Open("root:Antimaling2@@tcp(localhost:3306)/db_intern"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.AdminModel{}, &models.ManagerModel{}, &models.BarangModel{})

	return db
}
