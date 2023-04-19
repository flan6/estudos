package main

import (
	"image"
	"image/color"
	_ "image/png"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func main() {
	// Create a 100x100 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Fill the image with noise
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c := color.RGBA{
				uint8(x % 256),
				uint8(y % 256),
				uint8((x * y) % 256),
				255,
			}
			img.Set(x, y, c)
		}
	}

	// Create a WAV file with the image data
	file, err := os.Create("image.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a WAV encoder that writes to the file
	encoder := wav.NewEncoder(file, 44100, 16, 1, 1)
	defer encoder.Close()

	buffer := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  44100,
		},
		Data: make([]int, img.Bounds().Dx()*img.Bounds().Dy()),
	}

	i := 0
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c := img.At(x, y).(color.RGBA)
			// Map the color values to audio signals
			buffer.Data[i] = int(c.R)
			i++
		}
	}
}
