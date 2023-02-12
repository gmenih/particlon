package particlon

import (
	"fmt"
	"gmenih/particlon/src/particlon/base"
	"gmenih/particlon/src/particlon/particle"
	"gmenih/particlon/src/particlon/quad"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type ParticleSystem struct {
	sprites []*ebiten.Image
	debug   *ebiten.Image
	tree    *quad.Tree[*particle.Particle]

	size base.Vector

	state GameState
}

func NewParticleSystem(width, height float64) *ParticleSystem {
	return &ParticleSystem{
		debug: ebiten.NewImage(100, 100),
		tree:  quad.NewTree[*particle.Particle](quad.NewBounds(0, 0, width, height), nil),
		size:  base.VV(width, height),
		state: STATE_INIT,
	}
}

func (s *ParticleSystem) Update() error {
	switch s.state {
	case STATE_INIT:
		s.init()
	case STATE_PLAY:
		s.update()
	case STATE_PAUSE:
		s.init()
	}

	return nil
}

func (s *ParticleSystem) init() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.state = STATE_PLAY
	}
}

func (s *ParticleSystem) update() {
	s.tree.ForEach(func(p *particle.Particle) {
		neighbors := s.tree.QueryRange(quad.BB(p.Position, 20))
		for _, n := range neighbors {
			if n != p {
				p.Attract(n)
			}
		}

		p.Update(s.size.X, s.size.Y)
	})

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.state = STATE_PAUSE
	}

	s.tree.Rebalance()
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	s.debug.Fill(colornames.Black)
	ebitenutil.DebugPrintAt(s.debug, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()), 0, 0)
	ebitenutil.DebugPrintAt(s.debug, fmt.Sprintf("TPS: %0.0f", ebiten.ActualTPS()), 0, 20)

	s.tree.ForEach(func(p *particle.Particle) {
		p.Draw(screen)
	})
	// s.tree.Debug(screen, 0)

	screen.DrawImage(s.debug, nil)
}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(s.size.X), int(s.size.Y)
}
