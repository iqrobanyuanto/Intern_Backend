package middlewares

import (
	"Intern_Backend/utils/token"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v5"
	"github.com/didip/tollbooth/v5/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func RateLimitMiddleware(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.JSON(httpError.StatusCode, gin.H{"error": httpError.Message})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ManagerCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Setelah token divalidasi, kita dapat memeriksa role dari user yang terautentikasi.
		role, err := token.ExtractUserRole(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Jika role tidak sesuai, berikan pesan error dan hentikan proses.
		if role != "manager" && role != "admin" {
			c.String(http.StatusForbidden, "Access denied. Insufficient role.")
			c.Abort()
			return
		}
	}
}

func AdminCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Setelah token divalidasi, kita dapat memeriksa role dari user yang terautentikasi.
		role, err := token.ExtractUserRole(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Jika role tidak sesuai, berikan pesan error dan hentikan proses.
		if role != "admin" {
			c.String(http.StatusForbidden, "Access denied. Insufficient role.")
			c.Abort()
			return
		}
	}
}
