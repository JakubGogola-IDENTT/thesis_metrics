package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

var originalImagePath string
var replicatedImagePath string
var grayScale bool

func parseFlags() {
	flag.StringVar(&originalImagePath, "o", "", "path to original image")
	flag.StringVar(&replicatedImagePath, "r", "", "path to replicated image")
	flag.BoolVar(&grayScale, "g", false, "use gray scale")

	flag.Parse()
}

func readImage(pathToImage string) image.Image {
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(pathToImage)

	if err != nil {
		log.Fatal("Can't read image. Check if given path to file is correct.")
	}

	defer file.Close()

	imageData, _, err := image.Decode(file)

	if err != nil {
		log.Fatal("Can't decode image. Check if given file has correct format.")
	}

	return imageData
}

func compareColors(c1, c2 color.Color, isGrayScale bool) (diff float64) {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	if isGrayScale {
		return math.Abs(float64(r1 - r2))
	}

	diff += math.Abs(float64(r1 - r2))
	diff += math.Abs(float64(g1 - g2))
	diff += math.Abs(float64(b1 - b2))
	diff += math.Abs(float64(a1 - a2))

	return diff
}

func fitness(originalImage image.Image, replicatedImage image.Image) float64 {
	var score float64

	bounds := image.Pt(originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y)

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			score += compareColors(replicatedImage.At(x, y), originalImage.At(x, y), grayScale)
		}
	}

	return score
}

func main() {
	parseFlags()

	originalImage := readImage(originalImagePath)
	replicatedImage := readImage(replicatedImagePath)

	if originalImage.Bounds().Dx() != replicatedImage.Bounds().Dx() || originalImage.Bounds().Dy() != replicatedImage.Bounds().Dy() {
		log.Fatal("Images has different sizes")
	}

	score := fitness(originalImage, replicatedImage)

	fmt.Printf("Score: %f\n", score)
}
