package main

import (
	"fmt"
	"net/http"
)

// simple hello http func
func hellohttp(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello http!")
}

func main() {
	http.HandleFunc("/", hellohttp)
	//no addr will listen on every interface
	http.ListenAndServe(":9999", nil)
}
