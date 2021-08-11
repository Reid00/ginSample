package v3JsonResp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stu struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score,string"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// method1: return json string

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mesg": "this is first method, for 渲染Json, return json directly",
		})
	})

	r.GET("/json_struct", func(c *gin.Context) {
		stu := Stu{
			Name:  "Reid",
			Age:   22,
			Score: 99,
		}
		c.JSON(http.StatusOK, stu)
	})

	return r

}
