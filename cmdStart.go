package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

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

// CmdStart : subcommand to start a nucleome data server
func CmdStart(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	root := c.String("root")
	local := c.Bool("local")
	customCors := c.String("cors")

	corsOptions := nbdata.GetCors(customCors)
	mkdir(root)

	if nbdata.GuessURIType(uri) == "gsheet" {
		ctx := context.Background()
		config := nbdata.GsheetConfig()
		gA := asheets.NewGAgent(root)
		if !gA.HasCacheFile() {
			gA.GetClient(ctx, config)
		}
	}

	s := box.NewBox("NucleServer", VERSION).Port(port).CorsOptions(&corsOptions)
	router := s.GetRouter()

	idxRoot := filepath.Join(root, "index")
	mkdir(idxRoot)
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = nbdata.PluginTsv
	if uri != "" {
		l.Load(uri, router)
	}
	router.Use(nbdata.Cred)
	password := c.String("code")

	if password != "" {
		nbdata.InitCache(password)
		router.HandleFunc("/signin", nbdata.Signin)
		router.HandleFunc("/signout", nbdata.Signout)
		router.HandleFunc("/main.html", nbdata.MainHtml)
		router.Use(nbdata.SecureMiddleware)
		go s.Start("global")
	} else if local {
		go s.Start("local")
	} else {
		go s.Start("global")
	}

	// Graceful Shutdown
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	fmt.Println("Using Ctrl-C to Quit Program")
	select {
	case <-sigc:
	}
	s.Stop()

	log.Println("Exiting...")

	return nil
}
