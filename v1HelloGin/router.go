package v1HelloGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin, this is v1")

	})

	r.GET("/helloJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"resp": "hello gin, this is v1 json response",
		})
	})

	return r
}
