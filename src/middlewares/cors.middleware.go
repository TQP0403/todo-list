package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"X-Requested-With", "Content-Type", "Origin", "Authorization", "Accept", "Client-Security-Token", " Accept-Encoding", "x-access-token"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 24 * time.Hour
	config.AllowCredentials = true

	return cors.New(config)
}
