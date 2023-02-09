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

	sprite *ebiten.Image
}

func NewParticle(position, velocity base.Vector, kind ParticleKind, sprite *ebiten.Image) *Particle {
	return &Particle{
		Kind:     kind,
		Position: position,
		velocity: velocity,
		sprite:   sprite,
	}
}

func (p *Particle) Update() {
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

func (p *Particle) Attract(p2 *Particle) {
	diff := p.Position.Sub(p2.Position)
	distance := diff.Length()
	scale := 0.005

	if p.Kind != p2.Kind {
		scale -= 0.01
	} else if distance < 2 {
		scale *= -1
	}

	force := diff.Normalize().Scale(scale)
	p.velocity = p.velocity.Sub(force)
}
