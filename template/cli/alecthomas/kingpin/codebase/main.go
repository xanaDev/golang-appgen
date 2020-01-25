package main

import (
	"os"
	"{{ .AppName }}/cmd"
)


func main() {
	cmd.Commands(os.Args[1:])
}