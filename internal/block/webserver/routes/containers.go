package routes

import (
	"github.com/baswilson/block/internal/block"
	"github.com/gin-gonic/gin"
)

func Containers(c *gin.Context) {

	c.JSON(200, gin.H{
		"containers": block.Configs,
	})
}
