package v2Router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 从路由中获取参数

// 保存user
func UserSave(c *gin.Context) {
	username := c.Param("name")

	c.String(http.StatusOK, "Param method get: User: "+username+" saved already!")
}

// 通过query 的方法获取参数
func UserSaveByQuery(c *gin.Context) {
	username := c.Query("name")
	age := c.Query("age")

	c.String(http.StatusOK, "Query method get: User: "+username+" Age: "+age+"saved already!")
}
