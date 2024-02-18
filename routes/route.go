package routes

import (
	"github.com/didip/tollbooth/v5"
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

	rateLimiter := tollbooth.NewLimiter(10, nil)

	r.Use(middlewares.CorsMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", middlewares.RateLimitMiddleware(rateLimiter), controllers.Register)
	r.POST("/login-admin", middlewares.RateLimitMiddleware(rateLimiter), controllers.LoginAdmin)
	r.POST("/login-manager", middlewares.RateLimitMiddleware(rateLimiter), controllers.Login)

	managerMiddlewareRoute := r.Group("/get-product")
	managerMiddlewareRoute.Use(middlewares.ManagerCheckMiddleware())

	managerMiddlewareRoute.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	managerMiddlewareRoute.GET("/product", middlewares.RateLimitMiddleware(rateLimiter), controllers.GetByIdBarang)
	managerMiddlewareRoute.GET("/search", middlewares.RateLimitMiddleware(rateLimiter), controllers.SearchBarang)
	managerMiddlewareRoute.GET("/filter", middlewares.RateLimitMiddleware(rateLimiter), controllers.FilterBarang)

	adminMiddlewareRoute := r.Group("/update-product")
	adminMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())

	adminMiddlewareRoute.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	adminMiddlewareRoute.POST("/add", middlewares.RateLimitMiddleware(rateLimiter), controllers.Add)
	adminMiddlewareRoute.DELETE("/delete", middlewares.RateLimitMiddleware(rateLimiter), controllers.Delete)
	adminMiddlewareRoute.PUT("/update", middlewares.RateLimitMiddleware(rateLimiter), controllers.UpdateBarang)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
