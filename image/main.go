package main

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"os"
	"sync"
	"time"

	"libimage/shader"
)

func main() {
	const (
		width, height = 640, 360
		frames        = 100
	)

	logger := log.Default()
	logger.SetFlags(log.Ltime | log.Lshortfile)
	logger.SetPrefix("IMAGE: ")

	start := time.Now()
	imgs := []image.Image{}

	execTime := measureExecutionTime(func() {
		imgs = genImages(frames, width, height, shader.Circles)
	})
	logger.Println("genImages in", execTime)

	execTime = measureExecutionTime(func() {
		_, err := createGIF(imgs, "save/")
		if err != nil {
			logger.Fatal(err)
		}
	})
	logger.Println("createGIF in", execTime)

	logger.Println("Done in", time.Since(start))
}

func rgbaToPaletted(rgbaImg image.Image) *image.Paletted {
	palettedImg := image.NewPaletted(rgbaImg.Bounds(), palette.Plan9)
	draw.Draw(palettedImg, palettedImg.Rect, rgbaImg, rgbaImg.Bounds().Min, draw.Over)

	return palettedImg
}

func genImages(count, width, height int, fn shader.Shader) []image.Image {
	imgs := make([]image.Image, count)
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		i := i
		go func() {
			img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)

			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					img.Set(x, y, fn(image.Point{X: x, Y: y}, img.Bounds(), i))
				}
			}

			imgs[i] = img
			wg.Done()
		}()
	}

	wg.Wait()
	return imgs
}

func createGIF(imgs []image.Image, path string) (gif.GIF, error) {
	const (
		rate    = 70
		gifName = "animation.gif"
	)

	anim := gif.GIF{LoopCount: 0}
	anim.Image = make([]*image.Paletted, 0, len(imgs))

	for _, img := range imgs {
		anim.Image = append(anim.Image, img.(*image.Paletted))
		anim.Delay = append(anim.Delay, 1000/rate)
	}

	gitAnimation, err := os.Create(path + gifName)
	if err != nil {
		return anim, err
	}
	defer gitAnimation.Close()

	if err := gif.EncodeAll(gitAnimation, &anim); err != nil {
		return anim, err
	}

	return anim, nil
}

func measureExecutionTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
