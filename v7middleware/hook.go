package v7middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Set("name", "Reid") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值

		// 调用该请求的剩余处理程序
		c.Next()

		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("time cost: ", cost)
	}
}

// 全局路由使用
func UseHookPub() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	// gin.Default()默认使用了Logger和Recovery中间件，

	r := gin.New()

	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从hook 中间件中获取上下文

		log.Println("Name is: ", name)

		c.JSON(http.StatusOK, gin.H{
			"message": "hello" + name,
		})
	})

	return r
}

// 给/test2路由单独注册中间件（可注册多个）
func UseHookForSomeRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/test2", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	return r
}
