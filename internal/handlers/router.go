package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	app := gin.Default()
	app.Use(gin.Recovery())
	router := app.Group("/api" + "/v1")
	router.GET("/health-check", func(c *gin.Context) {
		c.Status(200)
	})
	return app
}
