package routes

import (
	"fmt"

	"github.com/baswilson/block/internal/block"
	"github.com/gin-gonic/gin"
)

type ActionDto struct {
	Action string `json:"action" binding:"required,oneof=start stop restart"`
	Config block.Config `json:"config" binding:"required,dive"`
}

func Action(c *gin.Context) {

	var body ActionDto

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	var err error
	switch body.Action {
		case "start":
			err = block.Run(&body.Config)
		case "stop":
			err =block.Stop(&body.Config)
		case "restart":
			err = block.Restart(&body.Config)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
