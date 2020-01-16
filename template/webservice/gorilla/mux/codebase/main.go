package main

import (
	"fmt"
	//"log"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening at :12345")
	fmt.Println(router)
	//http.ListenAndServe(":8080", router)
}
