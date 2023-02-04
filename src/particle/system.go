package particle

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const particleCount = 10_000

type ParticleSystem struct {
	Particles [particleCount]*Particle

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
	for i := range s.Particles {
		s.Particles[i].Update()
	}

	return nil
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	for i := range s.Particles {
		s.Particles[i].Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.Width, s.Height
}

func randomParticles(w, h float64) [particleCount]*Particle {
	var sprites = generateSprites()
	var particles [particleCount]*Particle

	for i := range particles {
		particles[i] = NewRandomParticle(w, h, sprites)
	}

	return particles
}
