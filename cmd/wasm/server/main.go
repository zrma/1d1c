package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "listen address")
		dir    = flag.String("dir", "static", "directory to serve")
	)

	filePath := filepath.Join(path(), *dir)

	flag.Parse()
	log.Println("listening on", *listen)
	log.Println("serving", filePath)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(filePath))))
}

func path() string {
	_, fileName, _, _ := runtime.Caller(1)
	return filepath.Dir(fileName)
}
