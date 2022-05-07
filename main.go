package main

import (
	"fmt"
	"log"
	"net/http"
)

// simple hello http func
func hellohttp(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello http!")
	log.Println("hello http")
}

func goodbyehttp(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "goodbye...")
	log.Println("goodbye http")
}

func main() {
	http.HandleFunc("/", hellohttp)

	http.HandleFunc("/goodbye/", goodbyehttp)

	// func as input value
	/*
		http.HandleFunc("/goodbye/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "goodbye...")
		log.Println("Goodbye http")
		})
	*/

	//no addr will listen on every interface
	http.ListenAndServe(":9999", nil)
}
