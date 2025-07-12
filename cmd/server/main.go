package main

import (
	"encoding/json"
	"log"
	"net/http"
	"thegame/pkg/player"

	"github.com/gorilla/websocket"
)

type state struct {
	players [2]player.Player
}

func main() {
	upgrader := websocket.Upgrader{}
	readychan := make(chan struct{})

	state := state{
		players: [2]player.Player{player.NewPlayer(), player.NewPlayer()},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error %s when upgrading connection to websocket", err)
			return
		}
		defer conn.Close()
		log.Printf("a player connected\n")

		var playeridx int
		select {
		case <-readychan:
			playeridx = 1
		default:
			playeridx = 0
			readychan <- struct{}{}
		}

		log.Printf("both players ready")

		for {
			for {
				msg, err := json.Marshal(state.players[1-playeridx])
				if err != nil {
					panic(err)
				}
				err = conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					panic(err)
				}

				log.Printf("send to %d", playeridx)

				_, msg, err = conn.ReadMessage()
				if err != nil {
					panic(err)
				}

				err = json.Unmarshal(msg, &state.players[playeridx])
				if err != nil {
					panic(err)
				}
			}
		}
	})
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe("192.168.101.34:4321", nil))
}
