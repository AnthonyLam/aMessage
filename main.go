package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mymux := http.NewServeMux()
	mymux.HandleFunc("/sms/new", pduToString)
	mymux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Println("Root dir")
	})
	mymux.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "bar\n")
	})

	// Configure the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mymux,
	}
	log.Fatal(server.ListenAndServe())
}

func pduToString(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You got a new message, but fuck you")
}
