package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg provided, got %d\n", len(os.Args)-1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()

	inputImage, err := jpeg.Decode(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	bounds := inputImage.Bounds().Bounds()
	outputImage := image.NewRGBA(bounds)
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			pixel := inputImage.At(i, j)
			pc := color.RGBAModel.Convert(pixel).(color.RGBA)

			grey := uint8(math.Round((float64(pc.R) + float64(pc.G) + float64(pc.B)) / 3))

			outputImage.Set(i, j, color.RGBA{
				R: grey,
				G: grey,
				B: grey,
				A: pc.A,
			})
		}
	}

	outputFile, err := os.Create("gray_output.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, outputImage, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
