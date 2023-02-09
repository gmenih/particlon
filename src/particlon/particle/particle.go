package particle

import (
	"gmenih/particlon/src/particlon/base"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleKind = uint8

type Particle struct {
	Kind ParticleKind

	position base.Vector
	velocity base.Vector

	sprite *ebiten.Image
}

func NewParticle(position, velocity base.Vector, kind ParticleKind, sprite *ebiten.Image) *Particle {
	return &Particle{
		Kind:     kind,
		position: position,
		velocity: velocity,
		sprite:   sprite,
	}
}

func (p *Particle) Update() {
	p.position = p.position.Add(p.velocity)
}

func (p *Particle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X-2, p.position.Y-2)
	screen.DrawImage(p.sprite, op)
}

func (p *Particle) Identity() (float64, float64) {
	return p.position.VV()
}
