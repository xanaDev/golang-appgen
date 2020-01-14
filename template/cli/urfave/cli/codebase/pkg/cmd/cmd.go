package cmd

import (
	"{{ .appname }}/pkg/consts"
	"github.com/urfave/cli"
	"os"
	"log"
)



//NewCmd as
func NewCmd()(app *cli.App){
	app = cli.NewApp()
	SetInfo(app)
	addCommands(app)
	return
}

//SetInfo Basic Info of cli
func SetInfo(app *cli.App){
	app.Name = consts.AppName
	app.Usage = consts.Usage
	app.Author = consts.Author
	app.Version = consts.Version
}

//Run to start accepting commands
func Run(app *cli.App){
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}
