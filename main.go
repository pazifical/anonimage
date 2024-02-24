package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

var testImagePath = "testdata/test_normal.png"
var testImageCheckPath = "testdata/test_normal_check.png"
var testShufflePath = "testdata/test_shuffle.png"
var testUnshufflePath = "testdata/test_unshuffle.png"

func main() {
	check()

	loadAndShuffle()
	loadAndUnshuffle()
	// loadAndShuffleTwice()
	// loadAndUnshuffleTwice()
}

func check() {
	img, err := loadImage(testImagePath)
	if err != nil {
		log.Fatal(err)
	}
	writer, err := os.Create(testImageCheckPath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, img)
	if err != nil {
		log.Fatal(err)
	}
}

func loadAndShuffle() {
	inputImage, err := loadImage(testImagePath)
	if err != nil {
		log.Fatal(err)
	}

	outputImage, err := shuffle(inputImage, "forwards")
	if err != nil {
		log.Fatal(err)
	}
	writer, err := os.Create(testShufflePath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, outputImage)
	if err != nil {
		log.Fatal(err)
	}
}

func loadAndShuffleTwice() {
	inputImage, err := loadImage(testImagePath)
	if err != nil {
		log.Fatal(err)
	}

	tempImage, err := shuffle(inputImage, "forwards")
	if err != nil {
		log.Fatal(err)
	}

	tempImage2 := rotate(tempImage)

	outputImage, err := shuffle(tempImage2, "forwards")
	if err != nil {
		log.Fatal(err)
	}

	writer, err := os.Create(testShufflePath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, outputImage)
	if err != nil {
		log.Fatal(err)
	}
}

func loadAndUnshuffle() {
	inputImage, err := loadImage(testShufflePath)
	if err != nil {
		log.Fatal(err)
	}

	outputImage, err := shuffle(inputImage, "backwards")
	if err != nil {
		log.Fatal(err)
	}
	writer, err := os.Create(testUnshufflePath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, outputImage)
	if err != nil {
		writer.Close()
		log.Fatal(err)
	}
}

func loadAndUnshuffleTwice() {
	inputImage, err := loadImage(testShufflePath)
	if err != nil {
		log.Fatal(err)
	}

	tempImage, err := shuffle(inputImage, "backwards")
	if err != nil {
		log.Fatal(err)
	}

	tempImage2 := rotate(tempImage)

	outputImage, err := shuffle(tempImage2, "backwards")
	if err != nil {
		log.Fatal(err)
	}

	writer, err := os.Create(testUnshufflePath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, outputImage)
	if err != nil {
		writer.Close()
		log.Fatal(err)
	}
}

func rotate(inputImage *image.RGBA) *image.RGBA {
	outputImage := image.NewRGBA(image.Rect(0, 0, inputImage.Bounds().Max.Y, inputImage.Bounds().Max.X))

	for x := range inputImage.Bounds().Max.X {
		for y := range inputImage.Bounds().Max.Y {
			outputImage.Set(y, x, inputImage.At(x, y))
		}
	}

	return outputImage
}

func writeBoundary(width, height int, inputImage, outputImage *image.RGBA) {
	for x := range width {
		outputImage.Set(x, 0, inputImage.At(x, 0))
		// outputImage.Set(x, height-1, inputImage.At(x, height-1))
	}
	for y := range height {
		outputImage.Set(0, y, inputImage.At(0, y))
		// outputImage.Set(width-1, y, inputImage.At(width-1, y))
	}
}

func shuffle(inputImage *image.RGBA, direction string) (*image.RGBA, error) {
	rect := image.Rect(0, 0, inputImage.Bounds().Max.X, inputImage.Bounds().Max.Y)
	outputImage := image.NewRGBA(rect)

	width := inputImage.Bounds().Max.X
	height := inputImage.Bounds().Max.Y

	writeBoundary(width, height, inputImage, outputImage)

	for y := 1; y < height; y++ {
		c := inputImage.At(0, y)
		r, g, b, _ := c.RGBA()
		xShift := int(r + g + b)

		switch direction {
		case "forwards":
			shiftForwards(y, xShift, width, inputImage, outputImage)
		case "backwards":
			shiftBackwards(y, xShift, width, inputImage, outputImage)
		}
	}

	return outputImage, nil
}

func shiftForwards(y, xShift, width int, inputImage, outputImage *image.RGBA) {
	for x := 1; x < width; x++ {
		newX := x + xShift
		if newX == 0 {
			panic("OH OH!")
		}
		for newX >= width {
			newX -= width
		}
		if newX < 0 {
			panic("OH OH!")
		}
		outputImage.Set(x, y, inputImage.At(newX, y))
	}
}
func shiftBackwards(y, xShift, width int, inputImage, outputImage *image.RGBA) {
	for x := 1; x < width-1; x++ {
		newX := x - xShift
		if newX == 0 {
			panic("OH OH!")
		}
		for newX < 0 {
			newX += width
		}
		if newX >= width {
			panic("OH OH!")
		}
		outputImage.Set(x, y, inputImage.At(newX, y))
	}
}

func decode(imageType string, f *os.File) (image.Image, error) {
	switch imageType {
	case "png":
		return png.Decode(f)
	case "jpeg":
		return jpeg.Decode(f)
	default:
		return image.NewRGBA(image.Rectangle{}), errors.New("not implemented")
	}
}

func loadImage(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, imageType, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	f.Seek(0, 0)
	loadedImage, err := decode(imageType, f)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(loadedImage.Bounds())
	for x := range loadedImage.Bounds().Max.X {
		for y := range loadedImage.Bounds().Max.Y {
			img.Set(x, y, loadedImage.At(x, y))
		}
	}

	return img, nil
}
