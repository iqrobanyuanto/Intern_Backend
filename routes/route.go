package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Intern_Backend/controllers"
	"Intern_Backend/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Use the CorsMiddleware here
	r.Use(middlewares.CorsMiddleware())

	r.POST("/register", controllers.Register)
	r.POST("/login-admin", controllers.LoginAdmin)
	r.POST("/login-manager", controllers.Login)

	managerMiddlewareRoute := r.Group("/get-product")
	managerMiddlewareRoute.Use(middlewares.ManagerCheckMiddleware())
	managerMiddlewareRoute.GET("/", controllers.GetByIdBarang)
	managerMiddlewareRoute.GET("/search", controllers.SearchBarang)
	managerMiddlewareRoute.GET("/filter", controllers.FilterBarang)

	adminMiddlewareRoute := r.Group("/update-product")
	adminMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())
	adminMiddlewareRoute.POST("/add", controllers.Add)
	adminMiddlewareRoute.DELETE("/delete", controllers.Delete)
	adminMiddlewareRoute.PUT("/update", controllers.UpdateBarang)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
