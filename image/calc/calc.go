package calc

import (
	"math"

	"libimage/vector"
)

func Step(tresh, value float64) float64 {
	if value < tresh {
		return 0.0
	} else {
		return 1.0
	}
}

func SmoothStep(min, max, value float64) float64 {
	if value < min {
		return 0.0
	} else if value > max {
		return 1.0
	} else {
		return value * value * value * (value*(value*6-15) + 10)
	}
}

func Fract2(vec vector.Vec2) vector.Vec2 {
	return vector.Vec2{
		X: vec.X - math.Floor(vec.X),
		Y: vec.Y - math.Floor(vec.Y),
	}
}

func Fract3(vec vector.Vec3) vector.Vec3 {
	res := vector.Vec3{}

	_, res.X = math.Modf(vec.X)
	_, res.Y = math.Modf(vec.Y)
	_, res.Z = math.Modf(vec.Z)

	return res
}
