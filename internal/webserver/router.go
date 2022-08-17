package webserver

import (
	"os"
	"time"

	"github.com/baswilson/block/internal/webserver/middleware"
	"github.com/baswilson/block/internal/webserver/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()
	router.Use(middleware.VerifyApiKey())
	router.SetTrustedProxies(nil)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/action", routes.Action)
	router.Run(":8080")
}
