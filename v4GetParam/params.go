package v4GetParam

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取query参数
func GetQueryParams(c *gin.Context) {
	username := c.DefaultQuery("username", "default Name:Reid")
	addr := c.Query("addr")

	// 输出json结果给调用发
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"addr":     addr,
	})

}

// 获取form 参数
func GetFormParam(c *gin.Context) {
	username := c.DefaultPostForm("username", "default PostFom: Reid")
	addr := c.PostForm("addr")

	c.JSON(http.StatusOK, gin.H{
		"mesg":     "ok",
		"username": username,
		"addr":     addr,
	})
}

// 获取Json参数
func GetJsonParam(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		fmt.Println("GetRawData failed.")
	}

	// 定义map 或者 struct 存储b 的value
	// var m map[string]interface{}
	type info struct {
		Msg  string `json:"msg"`
		Name string `json:"username"`
		Addr string `json:"addr"`
	}
	var m info
	// 反序列化
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("unmarshal failed, err: ", err)
	}

	c.JSON(http.StatusOK, m)

}

// 获取path参数
func GetPathParam(c *gin.Context) {
	username := c.Param("username")
	addr := c.Param("addr")

	c.JSON(http.StatusOK, gin.H{
		"mesg":     "ok",
		"username": username,
		"addr":     addr,
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/user", GetQueryParams)

	r.POST("/user/form", GetFormParam)

	r.POST("/user/json", GetJsonParam)

	r.GET("/user/:username/:addr", GetPathParam)

	return r
}
