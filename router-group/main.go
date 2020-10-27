package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "from %s", ctx.Request.URL)
}

func registerHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "from %s", ctx.Request.URL)
}

func main() {
	router := gin.Default()

	// 路由组
	v1 := router.Group("v1")
	{
		v1.POST("/login", loginHandler)
		v1.POST("/register", registerHandler)
	}

	v2 := router.Group("v2")
	{
		v2.POST("/login", loginHandler)
		v2.POST("/register", registerHandler)
	}

	_ = router.Run()
}
