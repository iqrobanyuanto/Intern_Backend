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
	username := utils.Getenv("DATABASE_USERNAME", "azureuser")
	password := utils.Getenv("DATABASE_PASSWORD", "^Bronya123")
	host := utils.Getenv("DATABASE_HOST", "mysqlserver3030.database.windows.net")
	port, converr := strconv.Atoi(utils.Getenv("DATABASE_PORT", "1433"))
	database := utils.Getenv("DATABASE_NAME", "industrial")

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
