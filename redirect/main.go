package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})

	router.GET("/test3", func(ctx *gin.Context) {
		// 路由重定向
		ctx.Request.URL.Path = "/test2"
		router.HandleContext(ctx)
	})

	router.GET("/test2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "200",
		})
	})

	_ = router.Run()
}
