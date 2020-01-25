package main

import (
	"{{ .AppName }}/router"
	"goji.io"
	"fmt"
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
	fmt.Println("Server started")
	mux := goji.NewMux()
	router.RegisterRoutes(mux)
	
}
