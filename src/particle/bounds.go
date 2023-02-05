package particle

type Bounds struct {
	Min Vector
	Max Vector
}

func BB(minX, minY, maxX, maxY float64) Bounds {
	return Bounds{
		Min: VV(minX, minY),
		Max: VV(maxX, maxY),
	}
}

func Around(v Vector, r float64) Bounds {
	return Bounds{
		Min: VV(v.X-r, v.Y-r),
		Max: VV(v.X+r, v.Y+r),
	}
}

func (b Bounds) Contains(v Vector) bool {
	return v.X >= b.Min.X && v.X <= b.Max.X && v.Y >= b.Min.Y && v.Y <= b.Max.Y
}

func (b Bounds) Intersects(b2 Bounds) bool {
	return b.Min.X <= b2.Max.X && b.Max.X >= b2.Min.X && b.Min.Y <= b2.Max.Y && b.Max.Y >= b2.Min.Y
}
