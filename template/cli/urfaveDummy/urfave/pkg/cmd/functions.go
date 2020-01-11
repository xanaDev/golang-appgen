package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"urfave/pkg/os"
	"urfave/pkg/arithmatic"
	"urfave/pkg/browser"
	"urfave/pkg/environment"
)


func getOsDetails(ctx *cli.Context){
	fmt.Println(os.GetOsDetails(cpu,name))
}

func addNumber(ctx *cli.Context){
	fmt.Println(arithmatic.Add(10,20))
}

func subtractNumber(ctx *cli.Context){
	fmt.Println(arithmatic.Subtract(10,20))
}
func divideNumber(ctx *cli.Context){
	fmt.Println(arithmatic.Division(10,20))
}
func multiplyNumber(ctx *cli.Context){
	fmt.Println(arithmatic.Multiply(10,20))
}

func openBrowser(ctx *cli.Context){
	fmt.Println(browser.Open("www.google.com"))
}

func getEnv(ctx *cli.Context){
	fmt.Println(environment.Get())
}
