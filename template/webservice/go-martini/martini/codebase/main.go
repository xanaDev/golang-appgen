package main

import (
	"codebase/router"
	"github.com/go-martini/martini"
	"codebase/forms"
	
)

func main() {

	m := martini.Classic()
	router.RegisterRoutes(m)
	
}
