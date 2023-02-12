package particlon

import (
	"gmenih/particlon/src/particlon/base"
	"gmenih/particlon/src/particlon/particle"
	"math/rand"
)

func randFloat(f, t float64) float64 {
	return f + rand.Float64()*(t-f)
}

func (s *ParticleSystem) InitParticles() {
	s.sprites = particle.GenerateSprites()
	w, h := s.size.VV()

	for i := 0; i < 6_000; i++ {
		r := uint8(rand.Intn(len(s.sprites)))
		p := particle.NewParticle(
			base.VV(randFloat(0, w), randFloat(0, h)),
			base.VV(0, 0),
			r,
			s.sprites[r],
		)

		s.tree.Insert(p)
	}
}
