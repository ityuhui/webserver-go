package main

import (
	"fmt"
	"net/http"
)

func api1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the api1!")
	fmt.Println("Endpoint Hit: api1")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html/")))
	http.HandleFunc("/api1", api1)
	http.ListenAndServe(":8080", nil)
}
