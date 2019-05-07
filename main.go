package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wbsCon *websocket.Conn
		err    error
		data   []byte
	)

	if wbsCon, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	go func() {
		var (
			err error
		)
		for {
			if err = wbsCon.WriteMessage(websocket.TextMessage, []byte("Heartbeat from server")); err != nil {
				return
			}
			time.Sleep(10 * time.Second)
		}
	}()

	for {
		if _, data, err = wbsCon.ReadMessage(); err != nil {
			goto ERR
		}
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	wbsCon.Close()

}

func api1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the api1!")
	fmt.Println("Endpoint Hit: api1")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html/")))
	http.HandleFunc("/api1", api1)
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8080", nil)
}
