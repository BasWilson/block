package routes

import (
	"fmt"

	"github.com/baswilson/block/internal/block"
	"github.com/gin-gonic/gin"
)

func Logs(c *gin.Context) {

	containerName := c.Param("container")
	output, err := block.Log(containerName)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"logs": output,
	})
}
