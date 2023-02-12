package particle

import (
	"gmenih/particlon/src/particlon/base"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleKind = uint8

type Particle struct {
	Kind     ParticleKind
	Position base.Vector

	velocity base.Vector
	size     float64

	sprite *ebiten.Image
}

func NewParticle(position, velocity base.Vector, kind ParticleKind, sprite *ebiten.Image) *Particle {
	return &Particle{
		Kind:     kind,
		Position: position,
		velocity: velocity,
		sprite:   sprite,
		size:     3.0,
	}
}

func (p *Particle) Update(w, h float64) {
	p.Position = p.Position.Add(p.velocity)
}

func (p *Particle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X-2, p.Position.Y-2)
	screen.DrawImage(p.sprite, op)
}

func (p *Particle) Identity() base.Vector {
	return p.Position
}
