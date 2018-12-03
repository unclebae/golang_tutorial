package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// var ws = new WebSocket("ws://localhost:3000/v1/ws")
	// ws.addEventListener("message", function(e) {console.log(e);});
	// ws.send("Hello kido")

	// : result
	// bubbles: false
	// cancelBubble: false
	// cancelable: false
	// composed: false
	// currentTarget: WebSocket {url: "ws://localhost:3000/v1/wc", readyState: 1, bufferedAmount: 0, onopen: null, onerror: null, …}
	// data: "Hello kido" <-- 결과
	// defaultPrevented: false
	// eventPhase: 0
	// isTrusted: true
	// lastEventId: ""
	// origin: "ws://localhost:3000"
	// path: []
	// ports: []
	// returnValue: true
	// source: null
	// srcElement: WebSocket {url: "ws://localhost:3000/v1/wc", readyState: 1, bufferedAmount: 0, onopen: null, onerror: null, …}
	// target: WebSocket {url: "ws://localhost:3000/v1/wc", readyState: 1, bufferedAmount: 0, onopen: null, onerror: null, …}
	// timeStamp: 159457.19999999984
	// type: "message"
	// __proto__: MessageEvent

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, _ := conn.ReadMessage()
				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	// var ws = new WebSocket("ws://localhost:3000/v2/ws")
	// ws.send("Hello Kido")

	// 결과 : 서버사이드에 결과 노출됨
	// Hello kido

	// ws.send(JSON.stringify({username: "KIDO"}))
	// 결과 : 서버사이드에 결과 노출됨
	// {"username":"KIDO"}
	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg))
			}
		}(conn)
	})

	// var ws = new WebSocket("ws://localhost:3000/v1/ws")
	// ws.addEventListener("message", function(e) {console.log(e);});
	// ws.readyState
	// ws.OK

	// ws.readyState
	// ws.CLOSE
	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch := time.Tick(10 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "KIDO",
					FirstName: "KIDO",
					LastName:  "BAE",
				})
			}
		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					conn.Close()
				}
			}
		}(conn)

		go func(conn *websocket.Conn) {
			ch := time.Tick(10 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "KIDO",
					FirstName: "KIDO",
					LastName:  "BAE",
				})
			}
		}(conn)
	})
	http.ListenAndServe(":3000", nil)
}

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
