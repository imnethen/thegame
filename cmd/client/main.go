package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func main() {
	ebiten.SetWindowSize(800, 600)
	game := Game{}

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
