package shader_test

import (
	"image"
	"testing"

	"libimage/shader"
)

func TestDefaultShader(t *testing.T) {
	res := shader.Circles(image.Point{500, 280}, image.Rectangle{Max: image.Point{600, 300}}, 1)
	t.Log(res)
	t.Fail()
}
