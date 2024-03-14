//
// To test this, start the program and then browse to http://localhost:8080/
//

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public
var webContent embed.FS
var webContentRoot, _ = fs.Sub(webContent, "public")

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(webContentRoot)))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
