package main

import (
	"encoding/json"
	"thegame/pkg/player"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	conn       *websocket.Conn
	thechannel chan player.Player

	yourPlayer  player.Player
	otherPlayer player.Player
}

func (g *Game) Update() error {
	g.yourPlayer.Update(true)
	g.otherPlayer.Update(false)
	select {
	case p := <-g.thechannel:
		g.otherPlayer = p
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.otherPlayer.Draw(screen)
	g.yourPlayer.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch := make(chan player.Player)

	ebiten.SetWindowSize(1600, 1200)
	game := Game{conn, ch, player.NewPlayer(), player.NewPlayer()}

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				panic(err)
			}

			var otherPlayer player.Player
			json.Unmarshal(msg, &otherPlayer)
			ch <- otherPlayer

			msg, err = json.Marshal(game.yourPlayer)
			if err != nil {
				panic(err)
			}
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				panic(err)
			}
		}
	}()

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
