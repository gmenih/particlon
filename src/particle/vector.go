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

func (v Vector) VV() (float64, float64) {
	return v.X, v.Y
}

func (v Vector) Copy() Vector {
	return Vector{
		X: v.X,
		Y: v.Y,
	}
}

func (v *Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector) Mag() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector) Distance(v2 Vector) float64 {
	return v.Sub(v2).Mag()
}

func (v Vector) Normalize() Vector {
	mag := v.Mag()
	return Vector{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

func (v Vector) Length() float64 {
	return v.Mag()
}

func (v Vector) Scale(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}
