package main

import (
	"os"
	"path/filepath"

	"github.com/nimezhu/nbdata"
	"github.com/urfave/cli"
)

const (
	VERSION = "0.1.2"
	DIR     = ".nucle"
)

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "Nucleome Data Server Tools"
	app.Usage = "nucleserver start -i [[google_sheet_id OR xls file]] -p [[port]]"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}
	home := nbdata.UserHomeDir()
	root := filepath.Join(home, DIR)
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "start a data server",
			Action: CmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data xls/google sheet id",
					Value: "",
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8611,
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory, default is $HOME/.nucle, in this directory store credentials and index files",
					Value: root,
				},
				cli.BoolFlag{
					Name:  "local,l",
					Usage: "serve 127.0.0.1 only",
				},
				cli.StringFlag{
					Name:  "code,c",
					Usage: "set password for server, override -l",
					Value: "",
				},
				cli.StringFlag{
					Name:  "cors",
					Usage: "add Customized CORS access",
					Value: "",
				},
			},
		},
		{
			Name:   "file",
			Usage:  "start a file server",
			Action: CmdFile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "root,r", //TODO
					Usage: "root directory",
					Value: home,
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8611,
				},
				cli.StringFlag{
					Name:  "cors",
					Usage: "add Customized CORS access",
					Value: "",
				},
			},
		},
	}
	app.Run(os.Args)
}
