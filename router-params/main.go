package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello %s", name)
	})

	// 此路由匹配 /user/mike/ 和 /user/mike/send
	// 如果没有其他路由匹配 /user/mike，他将重定向到 /user/mike/
	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
	})

	router.Run()
}
