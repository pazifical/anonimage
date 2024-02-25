package shuffle

import (
	"image/color"
)

type image1D struct {
	Pixels []color.Color
	Width  int
	Height int
}

func Process(inputPath, outputPath string, mode string) error {
	inputImage, err := loadFromFile(inputPath)
	if err != nil {
		return err
	}

	outputImage, err := shuffle(inputImage, mode)
	if err != nil {
		return err
	}

	err = writeToFile(outputImage, outputPath)
	if err != nil {
		return err
	}

	return nil
}
