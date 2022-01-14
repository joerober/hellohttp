package main

import (
	"log"
	"net/http"
)

// simple hello http func
func hellohttp(w http.ResponseWriter, _ *http.Request) {
	//fmt.Fprintf(w, "hello http!")
	log.Println("hello http")
}

func main() {
	http.HandleFunc("/", hellohttp)

	// func as input value
	http.HandleFunc("/goodbye/", func(w http.ResponseWriter, _ *http.Request) {
		//fmt.Fprintf(w, "goodbye...")
		log.Println("Goodbye http")
	})

	//no addr will listen on every interface
	http.ListenAndServe(":9999", nil)
}
