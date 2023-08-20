package vector

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Scalar(x float64) Vec2 {
	v.X *= x
	v.Y *= x

	return v
}

func (v Vec2) Sub(t Vec2) Vec2 {
	v.X -= t.X
	v.Y -= t.Y

	return v
}

func (v Vec2) Div(t Vec2) Vec2 {
	v.X = v.X / t.X
	v.Y = v.Y / t.Y

	return v
}

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Scalar(x float64) Vec3 {
	v.X *= x
	v.Y *= x
	v.Z *= x

	return v
}

func (v Vec3) Add(t Vec3) Vec3 {
	v.X += t.X
	v.Y += t.Y
	v.Z += t.Z

	return v
}

func (v Vec3) Multiply(t Vec3) float64 {
	return v.X*t.X + v.Y*t.Y + v.Z*t.Z
}
