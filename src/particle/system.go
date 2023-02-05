package particle

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const particleCount = 10_000

type ParticleSystem struct {
	Particles *QuadTree

	Width  int
	Height int
}

func NewParticleSystem(width, height int) *ParticleSystem {
	return &ParticleSystem{
		Width:     width,
		Height:    height,
		Particles: randomParticles(float64(width), float64(height)),
	}
}

func (s *ParticleSystem) Update() error {
	s.Particles.ForEach(func(p *Particle) {
		x, y := p.Position.VV()
		s.Particles.ForBoundary(BB(x-20, y-20, x+20, y+20), func(o *Particle) {
			p.Attract(o)
		})

		p.Update()
	})

	return nil
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	s.Particles.ForEach(func(p *Particle) {
		p.Draw(screen)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.Width, s.Height
}

func randomParticles(w, h float64) *QuadTree {
	var sprites = generateSprites()
	var particles *QuadTree = NewQuadTree(0, 0, w, h)

	for i := 0; i < particleCount; i++ {
		particles.Insert(NewRandomParticle(w, h, sprites))
	}

	return particles
}
