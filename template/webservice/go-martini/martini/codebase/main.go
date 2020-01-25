package main

import (
	"{{ .AppName }}/router"
	"github.com/go-martini/martini"
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
	
	m := martini.Classic()
	router.RegisterRoutes(m)
	
}
