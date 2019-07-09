package main

import (
	"context"
	"path"

	"github.com/gorilla/mux"
	"github.com/nimezhu/asheets"
	"github.com/nimezhu/box"
	"github.com/nimezhu/data"
	"github.com/nimezhu/nbdata"
	"github.com/urfave/cli"
)

func CmdStart(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	root := c.String("root")
	local := c.Bool("local")
	customCors := c.String("cors")

	corsOptions := nbdata.GetCors(customCors)

	router := mux.NewRouter()
	if nbdata.GuessURIType(uri) == "gsheet" {
		dir := path.Join(root, DIR)
		ctx := context.Background()
		config := nbdata.GsheetConfig()
		gA := asheets.NewGAgent(dir)
		if !gA.HasCacheFile() {
			gA.GetClient(ctx, config)
		}
	}
	s := box.NewBox("Nucleome Data Server", root, DIR, VERSION)
	s.InitRouter(router)
	s.InitHome(root)
	idxRoot := s.InitIdxRoot(root) //???
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = nbdata.PluginTsv
	if uri != "" {
		l.Load(uri, router)
	}
	router.Use(nbdata.Cred)
	password := c.String("code")
	if password != "" { //ADD PASSWORD CONTROL , MV IT TO WEB HTML
		nbdata.InitCache(password)
		router.HandleFunc("/signin", nbdata.Signin)
		router.HandleFunc("/signout", nbdata.Signout)
		router.HandleFunc("/main.html", nbdata.MainHtml)
		router.Use(nbdata.SecureMiddleware)
		s.StartDataServer(port, router, &corsOptions)
	} else if local {
		s.StartLocalServer(port, router, &corsOptions)
	} else {
		s.StartDataServer(port, router, &corsOptions)
	}

	return nil
}
