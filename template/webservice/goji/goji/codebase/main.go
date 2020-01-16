package main

import (
	"codebase/router"
	"goji.io"
	
)

func main() {
	mux := goji.NewMux()
	router.RegisterRoutes(mux)
	
}
