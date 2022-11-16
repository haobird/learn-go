package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	port = 8089
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", welcomeHandler)
	log.Printf("Server is running at %d port.\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	outputHTML(w, r, "static/index.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
