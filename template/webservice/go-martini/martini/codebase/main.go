package main

import (
	"codebase/router"
	"github.com/go-martini/martini"
	
	
)

func main() {

	m := martini.Classic()
	router.RegisterRoutes(m)
	
}
