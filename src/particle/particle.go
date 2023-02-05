package particle

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	Position Vector
	Velocity Vector
	Size     int
	Color    color.Color
	Sprite   *ebiten.Image
}

func NewParticle(position, velocity Vector, color color.Color, sprite *ebiten.Image, size int) *Particle {
	return &Particle{
		Position: position,
		Velocity: velocity,
		Size:     size,
		Color:    color,
		Sprite:   sprite,
	}
}

func NewRandomParticle(w, h float64, sprites Sprites) *Particle {
	x := rand.Float64() * w
	y := rand.Float64() * h
	c := colors[rand.Intn(len(colors))]

	v, ok := sprites[c]
	if !ok {
		panic("Fuck")
	}

	return NewParticle(
		VV(x, y),
		VV(0, 0),
		c,
		v,
		2,
	)
}

func (p *Particle) Update() {
	p.Position = p.Position.Add(p.Velocity)
}

func (p *Particle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Sprite, op)
}

func (p *Particle) Copy() *Particle {
	return &Particle{
		Position: p.Position,
		Velocity: p.Velocity,
		Color:    p.Color,
		Size:     p.Size,
	}
}

func (p *Particle) Attract(other *Particle) {
	p.Velocity = p.Velocity.Add(p.AttractForce(other))
}

func (p *Particle) AttractForce(other *Particle) Vector {
	if p == other {
		return VV(0, 0)
	}

	direction := other.Position.Sub(p.Position)
	distance := direction.Mag()
	direction = direction.Normalize()

	force := direction.Scale(0.1 / (distance * distance))

	return force
}
