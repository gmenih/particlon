package quad

type Bounds struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewBounds(minX, minY, maxX, maxY float64) Bounds {
	return Bounds{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.MinX && x <= b.MaxX && y >= b.MinY && y <= b.MaxY
}

func (b Bounds) Intersects(b2 Bounds) bool {
	return b.MinX <= b2.MaxX && b.MaxX >= b2.MinX && b.MinY <= b2.MaxY && b.MaxY >= b2.MinY
}
