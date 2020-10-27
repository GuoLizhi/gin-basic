package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "golang",
			"tag":  "<div>",
		}
		ctx.AsciiJSON(http.StatusOK, data)
	})
	_ = router.Run()
}
