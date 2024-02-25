package shuffle

import (
	"fmt"
	"image/color"
)

func calculateChunks(width, height int) int {
	pixelCount := width * height
	minChunks := int(pixelCount / 10)
	nChunks := minChunks
	for pixelCount%nChunks != 0 {
		nChunks += 1
	}
	fmt.Printf("INFO: Using %d chunks\n", nChunks)
	return nChunks
}

func shuffle(inputImage image1D, mode string) (image1D, error) {
	nChunks := calculateChunks(inputImage.Width, inputImage.Height)
	pixelCount := inputImage.Width * inputImage.Height
	chunkSize := pixelCount / nChunks

	startPositions := make([]int, nChunks)
	for i := range nChunks {
		startPositions[i] = i * chunkSize
	}

	outputImage := image1D{
		Pixels: make([]color.Color, pixelCount),
		Width:  inputImage.Width,
		Height: inputImage.Height,
	}

	for i := range len(outputImage.Pixels) {
		outputImage.Pixels[i] = color.RGBA{}
	}

	if mode == "shuffle" {
		i := 0
		offset := 0
		for i < pixelCount {
			for _, pos := range startPositions {
				outputImage.Pixels[pos+offset] = inputImage.Pixels[i]
				i += 1
			}
			offset += 1
		}
	} else if mode == "unshuffle" {
		i := 0
		offset := 0
		for i < pixelCount {
			for _, pos := range startPositions {
				outputImage.Pixels[i] = inputImage.Pixels[pos+offset]
				i += 1
			}
			offset += 1
		}
	}

	return outputImage, nil
}
