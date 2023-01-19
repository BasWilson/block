package routes

import (
	"github.com/gin-gonic/gin"
)

func Ready(c *gin.Context) {

	// Called by service to inform that it is ready to receive requests
	// this slave will then in turn inform the master that it is ready
	// the master will then call a webhook that is set in the config so the user can be notified

	c.JSON(200, gin.H{
		"ok": true,
	})
}
