package main

import (
	"codebase/router"
	"github.com/gin-gonic/gin"
	
)

func main() {

	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
