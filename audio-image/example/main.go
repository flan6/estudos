package main

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func main() {
	file, err := os.Open("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	file, err = os.Create("image.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const (
		sampleRate  = 44100
		bitDepth    = 16
		numChannels = 1
	)

	encoder := wav.NewEncoder(file, sampleRate, bitDepth, 1, 1)

	var frequencies []float64
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)

			r, g, b, _ := pixel.RGBA()
			frequency := math.Log(float64(r) + float64(g) + float64(b))
			frequencies = append(frequencies, frequency)
		}
	}

	buffer := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: numChannels,
			SampleRate:  sampleRate,
		},
		Data: make([]int, int(sampleRate)*numChannels),
	}

	for _, frequency := range frequencies {
		duration := 1.0 / sampleRate
		amplitude := math.MaxInt16 / 2
		phase := 0.0

		for i := range buffer.Data {
			buffer.Data[i] = int(float64(amplitude) * math.Sin(2*math.Pi*frequency*float64(i)*duration+phase))
		}

		err := encoder.Write(buffer)
		if err != nil {
			panic(err)
		}
	}

	err = encoder.Close()
	if err != nil {
		panic(err)
	}
}
