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

	for i := 0; i < 1000; i++ {
		r := uint8(rand.Intn(len(s.sprites)))
		p := particle.NewParticle(
			base.VV(randFloat(0, w), randFloat(0, h)),
			base.VV(0, 0),
			uint8(r),
			s.sprites[rand.Intn(len(s.sprites))],
		)

		s.tree.Insert(p)
	}
}
