package main

import (
	"go-asteroids-/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	roationPerSecond = math.Pi
	maxAcceleration  = 8.0
)

var currAcceleration float64

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector
	playerVelocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
		game:   game,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := roationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
	p.acceleration()
}
func (p *Player) acceleration() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if currAcceleration < maxAcceleration {
			currAcceleration = p.playerVelocity + 4.00
		}
		if currAcceleration >= maxAcceleration {
			currAcceleration = maxAcceleration
		}
		p.playerVelocity = currAcceleration

		dx := math.Cos(p.rotation) * -currAcceleration
		dy := math.Sin(p.rotation) * currAcceleration
		p.position.X += dx
		p.position.Y += dy
	}

}
