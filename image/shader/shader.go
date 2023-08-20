package shader

import (
	"image"
	"image/color"
	"math"

	"libimage/calc"
	"libimage/vector"
)

type Shader func(image.Point, image.Rectangle, int) color.Color

func Circles(coord image.Point, bounds image.Rectangle, curFrame int) color.Color {
	b := vector.Vec2{
		X: float64(bounds.Dx()),
		Y: float64(bounds.Dy()),
	}

	p := vector.Vec2{
		X: float64(coord.X),
		Y: b.Y - float64(coord.Y),
	}

	colors := vector.Vec3{}

	uv := p.Div(b).
		Sub(vector.Vec2{X: 0.5, Y: 0.5}).
		Scalar(2)
	uv.X *= b.X / b.Y
	uv0 := uv

	for i := 0.0; i < 4.0; i++ {
		uv = calc.Fract2(uv.Scalar(1.5)).
			Sub(vector.Vec2{X: 0.5, Y: 0.5})

		pal := palette(uv0.Length() + i*0.01 + float64(curFrame)*0.01)

		d := uv.Length() * math.Exp(-uv0.Length())
		d = math.Sin(d*8+float64(curFrame)*0.2) / 8
		d = math.Abs(d)

		colors = colors.Add(pal.Scalar(math.Pow(0.01/d, 1.01)))
	}

	return color.RGBA64{
		R: uint16(min(colors.X, 1.0) * 0xffff),
		G: uint16(min(colors.Y, 1.0) * 0xffff),
		B: uint16(min(colors.Z, 1.0) * 0xffff),
		A: uint16(1.0 * 0xffff),
	}
}

func palette(t float64) vector.Vec3 {
	a := [3]float64{0.5, 0.5, 0.5}
	b := [3]float64{0.5, 0.5, 0.5}
	c := [3]float64{1.0, 1.0, 1.0}
	d := [3]float64{0.263, 0.416, 0.557}
	var result [3]float64
	for i := range result {
		result[i] = a[i] + b[i]*math.Cos(6.28318*(c[i]*t+d[i]))
	}
	return vector.Vec3{
		X: result[0],
		Y: result[1],
		Z: result[2],
	}
}
