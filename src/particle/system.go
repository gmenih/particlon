package particle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ParticleSystem struct {
	Particles *QuadTree

	Width  int
	Height int

	sprites Sprites
}

func NewParticleSystem(width, height int) *ParticleSystem {
	return &ParticleSystem{
		Width:  width,
		Height: height,
		// Particles: NewQuadTree(0, 0, float64(width), float64(height)),
		Particles: randomParticles(float64(width), float64(height)),
		sprites:   generateSprites(),
	}
}

func (s *ParticleSystem) Update() error {
	s.Particles.ForEach(func(p *Particle) {
		s.Particles.ForBoundary(Around(p.Position, 50), func(o *Particle) {
			p.Attract(o)
		})

		p.Update()
	})

	s.Particles = s.Particles.Rebalance()

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		s.Particles.Insert(NewParticle(
			VV(float64(x), float64(y)),
			VV(0, 0),
			colors[0],
			s.sprites[colors[0]],
			3,
		))
	}

	return nil
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	s.Particles.ForEach(func(p *Particle) {
		p.Draw(screen)
	})

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
	// s.Particles.Debug(screen, 0)
}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.Width, s.Height
}

func randomParticles(w, h float64) *QuadTree {
	var sprites = generateSprites()
	var particles *QuadTree = NewQuadTree(0, 0, w, h, nil)

	for i := 0; i < PARTICLE_COUNT; i++ {
		particles.Insert(NewRandomParticle(w, h, sprites))
	}

	return particles
}
