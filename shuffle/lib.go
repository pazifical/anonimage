package shuffle

import (
	"image/color"
)

type Image1D struct {
	Pixels []color.Color
	Width  int
	Height int
}

func Process(inputPath, outputPath string) error {
	inputImage, err := loadImage1D(inputPath)
	if err != nil {
		return err
	}

	outputImage, err := shuffle(inputImage)
	if err != nil {
		return err
	}

	err = writeImage1D(outputImage, outputPath)
	if err != nil {
		return err
	}

	return nil
}
