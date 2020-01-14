package main

import (
	"os"
	"{{ .appname }}/cmd"
)


func main() {
	cmd.Commands(os.Args[1:])
}