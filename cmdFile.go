package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nimezhu/nbdata"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

var fileApp = nbdata.App{
	"CMU Fileome Server",
	"0.0.2",
}

func ls(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []os.FileInfo{}
	}
	return files
}

/* CNB
	File Server for Simple 3D Structure or Other files
  file server with meta.tsv for directory (alias)
	easily manage files with or without google sheets
	file server support range.

/* interface: ls dir and get/file */

//CmdFile : subcommand to start a simple file server
func CmdFile(c *cli.Context) error {
	root := c.String("root")
	port := c.Int("port")
	router := mux.NewRouter()
	//cors := data.CorsFactory(CORS)
	customCors := c.String("cors")

	corsOptions := nbdata.GetCors(customCors)

	//router.Use(cors)
	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		a, _ := json.Marshal(fileApp)
		w.Write(a)
	})
	router.HandleFunc("/ls/{dir:.*}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		dir := vars["dir"]
		fs := ls(path.Join(root, dir))
		fn := []string{}
		for _, f := range fs {
			fn = append(fn, f.Name())
		}
		j, err := json.Marshal(fn)
		if err == nil {
			w.Write(j)
		} else {
			w.Write([]byte("error"))
		}
	})
	router.PathPrefix("/get").Handler(http.StripPrefix("/get", http.FileServer(http.Dir(root))))
	hc := cors.New(corsOptions)
	handler := hc.Handler(router)
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: handler}
	log.Println("Please open http://127.0.0.1:" + strconv.Itoa(port))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	return err
}
