package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Intern_Backend/controllers"
	"Intern_Backend/middlewares"

	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.Use(middlewares.CorsMiddleware())

	r.POST("/register", controllers.Register)
	r.POST("/login-admin", controllers.LoginAdmin)
	r.POST("/login-manager", controllers.Login)

	managerMiddlewareRoute := r.Group("/get-product")
	managerMiddlewareRoute.Use(middlewares.ManagerCheckMiddleware())

	managerMiddlewareRoute.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	managerMiddlewareRoute.GET("/", controllers.GetByIdBarang)
	managerMiddlewareRoute.GET("/search", controllers.SearchBarang)
	managerMiddlewareRoute.GET("/filter", controllers.FilterBarang)

	adminMiddlewareRoute := r.Group("/update-product")
	adminMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())

	adminMiddlewareRoute.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	adminMiddlewareRoute.POST("/add", controllers.Add)
	adminMiddlewareRoute.DELETE("/delete", controllers.Delete)
	adminMiddlewareRoute.PUT("/update", controllers.UpdateBarang)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
