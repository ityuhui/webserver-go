package main

import (
	"fmt"
	"net/http"
)

// GetCheckInData : Get checkIn data from db
func GetCheckInData(w http.ResponseWriter, r *http.Request) {
	AllowCrossDomain(w)
	fmt.Fprintf(w, "1")
	fmt.Println("GET : /api/checkindata !")
}
