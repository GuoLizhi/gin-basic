package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// 比如请求http://localhost:8080?callback=callbackName
		// 会返回 callbackName({"foo":"bar"})
		ctx.JSONP(http.StatusOK, data)
	})
	_ = router.Run()
}
