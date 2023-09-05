package config

import (
	"Intern_Backend/models"
	"Intern_Backend/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.Getenv("DATABASE_USERNAME", "root")
	password := utils.Getenv("DATABASE_PASSWORD", "17ImJuuTmNYE7kbbBInA")
	host := utils.Getenv("DATABASE_HOST", "containers-us-west-43.railway.app")
	port := utils.Getenv("DATABASE_PORT", "6037")
	database := utils.Getenv("DATABASE_NAME", "railway")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.AdminModel{}, &models.ManagerModel{}, &models.BarangModel{})

	return db
}
