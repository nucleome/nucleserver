package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/nimezhu/asheets"
	"github.com/nimezhu/box"
	"github.com/nimezhu/data"
	"github.com/nimezhu/nbdata"
	"github.com/urfave/cli"
)

func mkdir(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}
}
func setExitSingal(s *box.Box) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			if sig == os.Interrupt || sig == syscall.SIGTERM {
				s.Stop()
				os.Exit(1)
			}
		}
	}()
}
func CmdStart(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	root := c.String("root")
	local := c.Bool("local")
	customCors := c.String("cors")

	corsOptions := nbdata.GetCors(customCors)
	// TODO Init Home and Directory
	mkdir(root)

	if nbdata.GuessURIType(uri) == "gsheet" {
		ctx := context.Background()
		config := nbdata.GsheetConfig()
		gA := asheets.NewGAgent(root)
		if !gA.HasCacheFile() {
			gA.GetClient(ctx, config)
		}
	}

	s := box.NewBox("Nucleome Data Server", VERSION).Port(port).CorsOptions(&corsOptions)
	router := s.GetRouter()
	setExitSingal(s) //TODO

	idxRoot := filepath.Join(root, "index")
	mkdir(idxRoot)
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = nbdata.PluginTsv
	if uri != "" {
		l.Load(uri, router)
	}
	router.Use(nbdata.Cred)
	password := c.String("code")
	fmt.Println("Using Ctrl-C to Quit Program") // STOP SHUT SIGNAL
	if password != "" {                         //ADD PASSWORD CONTROL , MV IT TO WEB HTML
		nbdata.InitCache(password)
		router.HandleFunc("/signin", nbdata.Signin)
		router.HandleFunc("/signout", nbdata.Signout)
		router.HandleFunc("/main.html", nbdata.MainHtml)
		router.Use(nbdata.SecureMiddleware)
		s.Start("global")
	} else if local {
		s.Start("local")
	} else {
		s.Start("global")
	}
	// Graceful s.Stop

	return nil
}
