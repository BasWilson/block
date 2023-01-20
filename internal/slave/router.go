package slave

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

	router.GET("/logs", logs)
	router.PUT("/config", applyConfig)
	router.Run(":8080")
}

type ApplyConfigDto struct {
	Config shared.Config `json:"config" binding:"required,dive"`
}

func applyConfig(c *gin.Context) {

	var body ApplyConfigDto

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var err = Restart(&body.Config)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func logs(c *gin.Context) {

	output, err := Log(Config.Settings.Name)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"logs": output,
	})
}
