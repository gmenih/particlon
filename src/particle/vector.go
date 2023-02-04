package particle

type Vector struct {
	X float64
	Y float64
}

func VV(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector) Sub(v2 *Vector) *Vector {
	return &Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}
