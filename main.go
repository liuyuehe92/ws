package main

import "golang.org/x/net/websocket"
import "github.com/codegangsta/cli"
import "strings"
import "fmt"
import "log"
import "os"

func main() {
	app := cli.NewApp()
	app.Name = "ws"
	app.Author = "haoxin"
	app.Version = "0.0.1"
	app.Usage = "💓 WebSocket CLI 💓"

	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			fmt.Println("invalid args")
			os.Exit(1)
		}

		url := c.Args()[0]

		ori := strings.Replace(url, "http", "ws", 1)
		ws, err := websocket.Dial(url, "", ori)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		for {
			var msg = make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", msg[:n])
		}
	}

	app.Run(os.Args)
}
