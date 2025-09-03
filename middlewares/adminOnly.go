package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminOnly checks if the user making the request is an admin.
// For now, we'll just simulate with a dummy check on the "X-User-Role" header.
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-User-Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "forbidden: only admins can perform this action",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
