package main

import (
	"fmt"
	"log"
	"net/http"
)

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)

	fmt.Fprintf(resp, "Search:%+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
