package main

import (
	"fmt"
	"net/http"
)

func api1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html/")))
	http.HandleFunc("/api1", api1)
	http.ListenAndServe(":8080", nil)
}
