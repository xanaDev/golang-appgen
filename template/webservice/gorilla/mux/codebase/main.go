package main

import (
	"fmt"
	//"log"
	"{{ .Logging.ImportPath}}"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening at :12345")
	fmt.Println(router)
	{{.Logging.Messages.Error}}
	//http.ListenAndServe(":8080", router)
}
