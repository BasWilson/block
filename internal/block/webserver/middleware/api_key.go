package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func VerifyApiKey() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// get key from header
		apiKey := "test"
		// apiKey := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")[1]

		// find the key in db, no result = 403
		if len(apiKey) == 0 {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		if os.Getenv("BLOCK_API_KEY") != apiKey {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		// add to context
		ctx.Set("api_key", apiKey)

		// finalize middleware
		ctx.Next()

		// do some logging in db, stuff after Next does not fuck up latency for user'
		// todo
	}
}
