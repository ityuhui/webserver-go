package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CheckInData struct {
	Already  int    `json:"already"`
	Left     int    `json:"left"`
	LastTime string `json:"lastTime"`
}

var cid CheckInData

func InitCheckInData() {
	cid = CheckInData{
		Already:  1,
		Left:     49,
		LastTime: "2021/5/11",
	}
}

func CheckInDataHandler(w http.ResponseWriter, r *http.Request) {
	AllowCrossDomain(w)
	if "GET" == r.Method {
		getCheckInData(w, r)
	} else if "POST" == r.Method {
		postCheckInData(w, r)
	} else if "DELETE" == r.Method {
		deleteCheckInData(w, r)
	}
}

func getCheckInData(w http.ResponseWriter, r *http.Request) {
	jsonCid, _ := json.Marshal(cid)
	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("GET : /api/checkindata !")
}

func postCheckInData(w http.ResponseWriter, r *http.Request) {
	cid.Already++
	cid.Left--
	cid.LastTime = time.Now().Format("2006/01/02")
	jsonCid, _ := json.Marshal(cid)

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("POST : /api/checkindata !")
}

func deleteCheckInData(w http.ResponseWriter, r *http.Request) {
	cid.Already--
	cid.Left++
	cid.LastTime = time.Now().Format("2006/01/02")
	jsonCid, _ := json.Marshal(cid)

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("DELETE : /api/checkindata !")
}
