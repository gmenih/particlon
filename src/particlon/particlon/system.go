package particlon

import (
	"gmenih/particlon/src/particlon/base"
	"gmenih/particlon/src/particlon/particle"
	"gmenih/particlon/src/particlon/quad"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ParticleSystem struct {
	sprites []*ebiten.Image
	tree    *quad.Tree[*particle.Particle]

	size base.Vector

	state GameState
}

func NewParticleSystem(width, height float64) *ParticleSystem {
	return &ParticleSystem{
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
		p.Update()
	})

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.state = STATE_PAUSE
	}

	s.tree.Rebalance()
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	// s.tree.Debug(screen, 0)
	s.tree.ForEach(func(p *particle.Particle) {
		p.Draw(screen)
	})

}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(s.size.X), int(s.size.Y)
}
