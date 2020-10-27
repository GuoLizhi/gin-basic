package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")

	_ = router.Run()
}
