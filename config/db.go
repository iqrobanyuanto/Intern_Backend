package config

import (
	"Intern_Backend/models"
	"Intern_Backend/utils"
	"fmt"
	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.Getenv("DBUSER", "root")
	password := utils.Getenv("DBPASS", "")
	host := utils.Getenv("DBHOST", "localhost")
	port, converr := strconv.Atoi(utils.Getenv("DBPORT", "3306"))
	database := utils.Getenv("DBNAME", "industrial")

	if converr != nil {
		panic(converr.Error())
	}

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", host, username, password, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.AdminModel{}, &models.ManagerModel{}, &models.BarangModel{})

	return db
}
