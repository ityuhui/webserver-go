package main

import (
	"fmt"
	"net/http"
	_ "time"

	"github.com/gorilla/websocket"
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

	/*go func() {
		var (
			err error
		)
		for {
			if err = wbsCon.WriteMessage(websocket.TextMessage, []byte("Websocket heartbeat from webserver-go")); err != nil {
				return
			}
			time.Sleep(10 * time.Second)
		}
	}()*/

	for {
		if _, data, err = wbsCon.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println("Received:")
		fmt.Println(data)
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
		fmt.Println("Sent:")
		fmt.Println(data)
	}

ERR:
	wbsCon.Close()

}

// AllowCrossDomain : Add http response header to allow cross domain
func AllowCrossDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	AllowCrossDomain(w)
	fmt.Fprintf(w, "Welcome to the /api/welcome !")
	fmt.Println("Endpoint Hit: /api/welcome")
}

func main() {
	fmt.Println("Webserver-go starts...")
	fmt.Println(" * static file serves at http://0.0.0.0:8080/")
	fmt.Println(" * RESTful API serves at http://0.0.0.0:8080/api/")
	fmt.Println(" * websocket serves at 0.0.0.0:8080/ws")
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/api/welcome", welcomeHandler)

	InitCheckInApp()
	http.HandleFunc("/api/checkindata", CheckInDataHandler)

	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8080", nil)
}
