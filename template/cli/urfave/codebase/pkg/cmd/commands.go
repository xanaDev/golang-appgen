package cmd

import (
	"github.com/urfave/cli"
)

var(
	name bool
	cpu bool
)

func addCommands(app *cli.App){
	app.Commands,_ = getCommands()
}



func getOsFlags()([]cli.Flag){
	return []cli.Flag{
        cli.BoolFlag{
            Name:        "name",
            Usage:       "will only show name",
            Destination: &name,
        },
        cli.BoolFlag{
            Name:        "cpu",
            Usage:       "shows no of cpu cores",
            Destination: &cpu,
        },
	}
}

func getCommands()([]cli.Command,error){
	return []cli.Command{
		{
			Name:"os",
			Aliases: []string{"system"},
			Usage: "get operating system details",
			Action: getOsDetails,
			Flags: getOsFlags(),
		},
		{
			Name:"add",
			Aliases: []string{"add"},
			Usage: "add two numbers",
			Action: addNumber,
		},
		{
			Name:"sub",
			Aliases: []string{"sub"},
			Usage: "subtract two numbers",
			Action: subtractNumber,
		},
		{
			Name:"div",
			Aliases: []string{"div"},
			Usage: "divides two number",
			Action: divideNumber,
		},
		{
			Name:"mul",
			Aliases: []string{"mul"},
			Usage: "multiply two numbers",
			Action: multiplyNumber,
		},
		{
			Name:"env",
			Aliases: []string{"e"},
			Usage: "gives envrionment variables list",
			Action: getEnv,
		},
		{
			Name:"open",
			Aliases: []string{"open"},
			Usage: "opens up a url in default browser",
			Action: openBrowser,
		},
	},nil
}

