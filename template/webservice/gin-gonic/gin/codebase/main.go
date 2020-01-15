package main

import (
	"{{ .appname }}/router"
	"github.com/gin-gonic/gin"
	
)

func main() {

	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
