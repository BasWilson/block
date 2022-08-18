package routes

import (
	"github.com/baswilson/block/internal/block"
	"github.com/gin-gonic/gin"
)

func Containers(c *gin.Context) {

	// create empty slice of containers
	containers := make([]block.Config, len(block.Configs))
	
	for i := 0; i < len(containers); i++ {
		container := block.Configs[i]
		for x := 0; x < len(container.Variables); x++ {
			container.Variables[x].Value = "*************";
		}
	}

	c.JSON(200, gin.H{
		"containers": containers,
	})
}
