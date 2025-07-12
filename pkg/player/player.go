package player

import (
	"image/color"
	"math"

	vec "github.com/deeean/go-vector/vector2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const speed = 0.3
const radius = 60

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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		p.Vel = *p.Pos.Sub(vec.New(float64(mouseX), float64(mouseY))).MulScalar(-1.).Normalize().MulScalar(p.Vel.Magnitude())
	}
}

func (p *Player) handleBorders(border vec.Vector2) {
	if p.Pos.X-radius < 0 {
		p.Pos.X = radius
		p.Vel.X = math.Abs(p.Vel.X)
	}
	if p.Pos.X+radius > border.X {
		p.Pos.X = border.X - radius
		p.Vel.X = -math.Abs(p.Vel.X)
	}
	if p.Pos.Y-radius < 0 {
		p.Pos.Y = radius
		p.Vel.Y = math.Abs(p.Vel.Y)
	}
	if p.Pos.Y+radius > border.Y {
		p.Pos.Y = border.Y - radius
		p.Vel.Y = -math.Abs(p.Vel.Y)
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
	p.handleBorders(*vec.New(1600, 1200))
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.Pos.X), float32(p.Pos.Y), radius, color.White, false)
}
