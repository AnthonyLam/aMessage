package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Set the port here")
	flag.Parse()

	mymux := http.NewServeMux()
	mymux.HandleFunc("/sms/new", SmsToObjAndBack)
	//mymux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Println("Root dir")
	//})

	// Configure the server
	server := &http.Server{
		Addr:    ":" + *port,
		Handler: mymux,
	}
	fmt.Println(*port)
	log.Fatal(server.ListenAndServe())
}

func SmsToObjAndBack(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Println("Got a POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Read bytes
		dat, err1 := ioutil.ReadAll(req.Body)
		if err1 != nil {
			fmt.Println("Too many bytes!!! AAHHH Vampires!")
			return
		}

		message := NewSms(dat)
		fmt.Println(message.String())
		_, err := w.Write(message.Bytes())
		if err != nil {
			fmt.Println("Couldn't marshal bytes")
		}
	} else {
		fmt.Println("Got other method")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Response")
}
