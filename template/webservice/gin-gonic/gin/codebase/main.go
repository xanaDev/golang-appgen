package main

import (
	"{{ .AppName }}/router"
	"github.com/gin-gonic/gin"
	{{ if .Logging.ImportPath }}
	_"{{ .AppName }}/logger"
	"{{ .Logging.ImportPath }}"
	{{end}}
	
)

func main() {

	{{ .Logging.Messages.Info }}
	{{ .Logging.Messages.Error }}
	{{ .Logging.Messages.Warn }}
	{{ .Logging.Messages.Debug }}
	
	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
