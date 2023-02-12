package particle

var forces = [6][6]float64{
	{1, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0},
	{0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 1},
}

func (p *Particle) getScale(p2 *Particle) float64 {
	return forces[p.Kind][p2.Kind]
}

func (p *Particle) Attract(p2 *Particle) {
	distance := p.Position.Distance(p2.Position)
	if distance < (2*p.size)+1 {
		// Move both points distance / 2 apart
		p.Position = p.Position.Add(p.Position.Sub(p2.Position).Normalize().Scale(distance / 2))
		p2.Position = p2.Position.Add(p2.Position.Sub(p.Position).Normalize().Scale(distance / 2))
	} else {
		// p.velocity = p.velocity.Add(p2.Position.Sub(p.Position).Normalize().Scale(scale / distance))
		// p2.velocity = p2.velocity.Add(p.Position.Sub(p2.Position).Normalize().Scale(scale / distance))
	}
}
