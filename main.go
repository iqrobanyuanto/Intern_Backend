package main

import (
	"Intern_Backend/config"
	"Intern_Backend/docs"
	"Intern_Backend/routes"
)

func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Test API"
	docs.SwaggerInfo.Description = "Testing API Produk."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
