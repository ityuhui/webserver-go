package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type checkInData struct {
	Already int      `json:"already"`
	Left    int      `json:"left"`
	History []string `json:"history"`
}

type checkInAppInstance struct {
	db *sql.DB
}

func NewAppInstance() *checkInAppInstance {
	appInstance := &checkInAppInstance{
		db: nil,
	}
	return appInstance
}

var app *checkInAppInstance = nil

func InitCheckInApp() {
	app = NewAppInstance()
	initDB()
}

func checkDBErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	dbName := os.Getenv("CIDB_FILE")
	if dbName == "" {
		dbName = "./bkbci.db"
	}
	db, err := sql.Open("sqlite3", "dbName")
	checkDBErr(err)
	createTableIfNeed()
	app.db = db
}

func createTableIfNeed() {
	sql_table := `CREATE TABLE IF NOT EXISTS "cid" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"already" INTEGER NULL,
		"left" INTEGER NULL
	);
	
	CREATE TABLE IF NOT EXISTS "cih" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"date" TEXT NULL
	);`

	app.db.Exec(sql_table)
}

func getCidTable() (int, int) {
	cidRows, err := app.db.Query("SELECT already,left FROM cid")
	checkDBErr(err)
	var already int
	var left int
	for cidRows.Next() {
		err = cidRows.Scan(&already, &left)
		checkDBErr(err)
		fmt.Println(already)
		fmt.Println(left)
	}
	return already, left
}

func getCihTable() []string {
	cihRows, err := app.db.Query("SELECT date FROM cih")
	checkDBErr(err)
	var history []string
	for cihRows.Next() {
		var date string
		err = cihRows.Scan(&date)
		checkDBErr(err)
		fmt.Println(date)
		history = append(history, date)
	}
	return history
}

func getDataFromDB() checkInData {
	already, left := getCidTable()
	history := getCihTable()
	return checkInData{
		Already: already,
		Left:    left,
		History: history,
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
	jsonCid, _ := json.Marshal(getDataFromDB())
	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("GET : /api/checkindata !")
}

func postCheckInData(w http.ResponseWriter, r *http.Request) {
	incCidTable()
	incCihTable()
	jsonCid, _ := json.Marshal(getDataFromDB())

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("POST : /api/checkindata !")
}

func deleteCheckInData(w http.ResponseWriter, r *http.Request) {
	decCidTable()
	decCihTable()
	jsonCid, _ := json.Marshal(getDataFromDB())

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("DELETE : /api/checkindata !")
}

func incCidTable() {

}

func incCihTable() {

}

func decCidTable() {

}

func decCihTable() {

}
