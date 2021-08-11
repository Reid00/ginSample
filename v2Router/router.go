package v2Router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 封装好回调函数
func retv2RouterMethod(c *gin.Context) {
	c.String(http.StatusOK, "hello gin, "+strings.ToLower(c.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	root := r.Group("/")
	{
		root.GET("", retv2RouterMethod)

		root.POST("", retv2RouterMethod)

		root.PUT("", retv2RouterMethod)

		root.DELETE("", retv2RouterMethod)

		// // or replace up with below
		// root.Any("", retv2RouterMethod)

	}

	// add router user
	user := r.Group("/user")
	{
		user.GET("/:name", UserSave)
		user.GET("", UserSaveByQuery)
	}

	return r
}
