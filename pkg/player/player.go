package player

import (
	"image/color"

	vec "github.com/deeean/go-vector/vector2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const speed = 0.3

type Player struct {
	Pos vec.Vector2 `json:"pos"`
	Vel vec.Vector2 `json:"vel"`
}

func NewPlayer() Player {
	return Player{
		Pos: *vec.New(0., 0.),
		Vel: *vec.New(0., 0.),
	}
}

func (p *Player) handleInput(acc *vec.Vector2) {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		acc.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		acc.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		acc.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		acc.Y += 1
	}
}

func (p *Player) Update(controlled bool) {
	acc := vec.New(0., 0.)

	if controlled {
		p.handleInput(acc)
	}
	acc = acc.Normalize().MulScalar(speed)

	p.Vel = *p.Vel.Add(acc)
	p.Pos = *p.Pos.Add(&p.Vel)
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.Pos.X), float32(p.Pos.Y), 60, color.White, false)
}
