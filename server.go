package aMessage

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Set the port here")
	flag.Parse()

	mymux := http.NewServeMux()
	mymux.HandleFunc("/sms/new", pduToString)
	mymux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Root dir")
	})

	// Configure the server
	server := &http.Server{
		Addr:    ":" + *port,
		Handler: mymux,
	}
	fmt.Println(*port)
	log.Fatal(server.ListenAndServe())
}

func pduToString(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You got a new message, but fuck you")
}
