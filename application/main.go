package main

import (
	"github.com/Jadepypy/distributed-social-media/application/internal"
	"github.com/Jadepypy/distributed-social-media/application/internal/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// TODO: add cookie settings
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("dsm", store))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	l := middleware.LogInMiddlewareBuilder{}
	r.Use(l.CheckLogIn())
	u := internal.NewUserHandler()
	u.RegisterRoutes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
