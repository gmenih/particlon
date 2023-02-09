package quad

import "gmenih/particlon/src/particlon/base"

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

func (b Bounds) Contains(v base.Vector) bool {
	return v.X >= b.MinX && v.X <= b.MaxX && v.Y >= b.MinY && v.Y <= b.MaxY
}

func (b Bounds) Intersects(b2 Bounds) bool {
	return b.MinX <= b2.MaxX && b.MaxX >= b2.MinX && b.MinY <= b2.MaxY && b.MaxY >= b2.MinY
}

func BB(v base.Vector, area float64) Bounds {
	return NewBounds(v.X-area, v.Y-area, v.X+area, v.Y+area)
}
