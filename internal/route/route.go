package route

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		// "Hello, World!" 반환
		c.String(200, "Hello, World!")
	})

	router.GET("/hello", func(c *gin.Context) {
		// "Hello, World!" 반환
		c.String(200, "Hello, World!")
	})

	return router
}
