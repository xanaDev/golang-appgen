package main

import (
	"codebase/router"
	"goji.io"
	"fmt"
	
)

func main() {
	fmt.Println("Server started")
	mux := goji.NewMux()
	router.RegisterRoutes(mux)
	
}
