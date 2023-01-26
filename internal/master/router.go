package master

import (
	"fmt"
	"os"
	"time"

	"github.com/baswilson/block/internal/shared"
	"github.com/baswilson/block/internal/shared/middleware"
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

	router.PUT("/config", putConfig)
	router.Run(":8080")
}

type PutConfigDto struct {
	Config shared.Config `json:"config" binding:"required,dive"`
}

func putConfig(c *gin.Context) {

	var body PutConfigDto

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	AddConfig(&body.Config)

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
