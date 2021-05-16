package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type checkInData struct {
	Already int    `json:"already"`
	Left    int    `json:"left"`
	Date    string `json:"date"`
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
	db, err := sql.Open("sqlite3", dbName)
	checkDBErr(err)
	app.db = db
	createTableIfNeed()
	initTable()
}

func createTableIfNeed() {
	sql_table := `CREATE TABLE IF NOT EXISTS "cid" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"already" INTEGER NULL,
		"left" INTEGER NULL,
		"date" TEXT NULL
	);`

	app.db.Exec(sql_table)
}

func initTable() {
	cidRows, err := app.db.Query("SELECT count(*) FROM cid")
	checkDBErr(err)
	var count int
	for cidRows.Next() {
		err = cidRows.Scan(&count)
		checkDBErr(err)
	}
	if count == 0 {
		insertTheDefaultRecord()
	}
}

func insertTheDefaultRecord() {
	sql := `INSERT INTO cid 
		("already", "left", "date" )
		VALUES 
		(	0,		5,		"" )`

	app.db.Exec(sql)
}

func getLastAlreadyAndLeft() (int, int) {
	cidRows, err := app.db.Query("SELECT already,left FROM cid order by id desc limit 1")
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

func getDataFromDB() []*checkInData {
	var result []*checkInData
	cidRows, err := app.db.Query("SELECT already,left,date FROM cid")
	checkDBErr(err)
	for cidRows.Next() {
		var already int
		var left int
		var date string
		err = cidRows.Scan(&already, &left, &date)
		checkDBErr(err)
		fmt.Println(already)
		fmt.Println(left)
		fmt.Println(date)
		cid := &checkInData{
			Already: already,
			Left:    left,
			Date:    date,
		}
		result = append(result, cid)
	}
	return result
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
	incCid()
	jsonCid, _ := json.Marshal(getDataFromDB())

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("POST : /api/checkindata !")
}

func deleteCheckInData(w http.ResponseWriter, r *http.Request) {
	decCid()
	jsonCid, _ := json.Marshal(getDataFromDB())

	fmt.Fprintf(w, string(jsonCid))
	fmt.Println("DELETE : /api/checkindata !")
}

func incCid() {
	already, left := getLastAlreadyAndLeft()

	stmt, err := app.db.Prepare("INSERT INTO cid (already, left, date) values(?,?,?)")
	checkDBErr(err)

	res, err := stmt.Exec(already+1, left-1, time.Now().Format("2006-01-02"))
	checkDBErr(err)

	id, err := res.LastInsertId()
	checkDBErr(err)

	fmt.Println(id)
}

func decCid() {
	sql_delete_last_item := "delete from cid where id in (select id from cid order by id desc limit 1)"

	res, err := app.db.Exec(sql_delete_last_item)
	checkDBErr(err)

	affect, err := res.RowsAffected()
	checkDBErr(err)

	fmt.Println(affect)
}
