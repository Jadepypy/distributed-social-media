package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogInMiddlewareBuilder struct{}

func (m *LogInMiddlewareBuilder) CheckLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/user/login" || path == "/user/signup" {
			c.Next()
			return
		}

		sess := sessions.Default(c)
		userID := sess.Get("user_id")
		if userID == nil {
			c.JSON(401, gin.H{"message": "Not logged in"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
