package main

import (
	"Intern_Backend/config"
	"Intern_Backend/docs"
	"Intern_Backend/routes"
	"Intern_Backend/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Test API"
	docs.SwaggerInfo.Description = "Testing API Produk."
	docs.SwaggerInfo.Version = "1.0"
	//could be changed, based on the services domain
	docs.SwaggerInfo.Host = "industrialapi.azurewebsites.net"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
