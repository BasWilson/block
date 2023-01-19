package routes

import (
	"fmt"

	"github.com/baswilson/block/internal/shared"
	"github.com/baswilson/block/internal/slave"
	"github.com/gin-gonic/gin"
)

type ActionDto struct {
	Action string        `json:"action" binding:"required,oneof=start stop restart"`
	Config shared.Config `json:"config" binding:"required,dive"`
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
		err = slave.Run(&body.Config)
	case "stop":
		err = slave.Stop(&body.Config)
	case "restart":
		err = slave.Restart(&body.Config)
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
