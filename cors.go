package medutils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CorsMiddleware() gin.HandlerFunc {
	medCors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Set-Cookie", "X-PINGOTHER", "x-auth-token", "x-refresh-token", "csrf-key"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	return medCors
}
