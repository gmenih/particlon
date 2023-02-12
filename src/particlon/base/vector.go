package base

import "math"

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

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Distance(v2 Vector) float64 {
	return v.Sub(v2).Length()
}

func (v Vector) Limit(max float64) Vector {
	l := v.Length()
	if l > max {
		return v.Normalize().Scale(max)
	}
	return v
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{
		X: v.X / l,
		Y: v.Y / l,
	}
}

func (v Vector) Scale(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}
