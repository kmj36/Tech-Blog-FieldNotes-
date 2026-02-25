package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/configs"
)

func CORSHandler(cfg *configs.Config) gin.HandlerFunc {
	var allowOrigins []string

	if gin.Mode() == gin.ReleaseMode {
		allowOrigins = []string{cfg.ProductionDomain}
	} else {
		allowOrigins = []string{cfg.TestDomain}
	}

	return cors.New(cors.Config{
        AllowOrigins:     allowOrigins,
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}