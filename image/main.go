package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"golang.org/x/exp/constraints"
)

func main() {
	const (
		width, height = 300, 300
		frames        = 100
	)

	rand.Seed(time.Now().UnixNano())
	logger := log.Default()
	logger.SetFlags(log.Ltime | log.Lshortfile)
	logger.SetPrefix("IMAGE: ")

	anim := gif.GIF{LoopCount: 1}

	for i := 0; i < frames; i++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				scaledY := (y) * (255) / (height)
				scaledX := (x) * (255) / (width)

				r := uint8(scaledX)
				g := uint8(min(math.Pow(float64(scaledY), 0.999), 255))
				b := uint8(min(math.Pow(float64(scaledX+scaledY), 0.9), 255))

				img.Set(x, y, color.NRGBA{r, g, b, uint8(255)})
			}
		}

		scaledI := (i) * (255) / (frames)

		for x := -50; x < 50; x++ {
			for y := -50; y < 50; y++ {
				if x*x+y*y < 50*50 {
					img.Set(width/2+x, height/2+y, color.NRGBA{
						R: uint8(scaledI),
						G: uint8(scaledI),
						B: uint8(255),
						A: 255,
					})
				}
			}
		}

		anim.Image = append(anim.Image, rgbaToPaletted(img))
		anim.Delay = append(anim.Delay, 1)

		pngIMG, err := os.Create(fmt.Sprintf("imgs/image_%d.png", i))
		if err != nil {
			logger.Fatal(err)
		}

		err = png.Encode(pngIMG, img)
		if err != nil {
			logger.Fatal(err)
		}
		pngIMG.Close()
	}

	gitAnimation, err := os.Create("animation.gif")
	if err != nil {
		logger.Fatal(err)
	}
	defer gitAnimation.Close()

	if err := gif.EncodeAll(gitAnimation, &anim); err != nil {
		logger.Fatal(err)
	}

	logger.Println("Done")
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}

func rgbaToPaletted(rgbaImg *image.RGBA) *image.Paletted {
	palettedImg := image.NewPaletted(rgbaImg.Bounds(), palette.Plan9)
	draw.Draw(palettedImg, palettedImg.Rect, rgbaImg, rgbaImg.Bounds().Min, draw.Over)

	return palettedImg
}
