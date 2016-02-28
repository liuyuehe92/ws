package main

import "golang.org/x/net/websocket"
import "net/http"
import "time"
import "fmt"

func main() {
	wsHandler := websocket.Handler(func(ws *websocket.Conn) {
		for {
			time.Sleep(1 * time.Second)
			ws.Write([]byte("hello"))
		}
	})

	fmt.Println("serve on 3000")
	http.ListenAndServe(":3000", wsHandler)
}
