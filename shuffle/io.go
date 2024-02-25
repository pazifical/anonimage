package shuffle

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func loadFromFile(imagePath string) (image1D, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return image1D{}, fmt.Errorf("opening image '%s': %w", imagePath, err)
	}
	defer f.Close()

	_, imageType, err := image.Decode(f)
	if err != nil {
		return image1D{}, fmt.Errorf("decoding image '%s': %w", imagePath, err)
	}

	f.Seek(0, 0)
	loadedImage, err := decode(imageType, f)
	if err != nil {
		return image1D{}, fmt.Errorf("loading image '%s': %w", imagePath, err)
	}

	width := loadedImage.Bounds().Max.X
	height := loadedImage.Bounds().Max.Y
	pixels := make([]color.Color, width*height)

	for y := range height {
		for x := range width {
			pixels[y*width+x] = loadedImage.At(x, y)
		}
	}

	fmt.Printf("INFO: Imported %dx%d image with %d pixels (%s)\n", width, height, len(pixels), imagePath)

	return image1D{
		Pixels: pixels,
		Width:  width,
		Height: height,
	}, nil
}

func writeToFile(image1d image1D, outputPath string) error {
	tempImage := image.NewRGBA(image.Rect(0, 0, image1d.Width, image1d.Height))

	for y := range image1d.Height {
		for x := range image1d.Width {
			tempImage.Set(x, y, image1d.Pixels[y*image1d.Width+x])
		}
	}

	writer, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating '%s': %w", outputPath, err)
	}
	defer writer.Close()

	err = png.Encode(writer, tempImage)
	if err != nil {
		writer.Close()
		return fmt.Errorf("writing '%s': %w", outputPath, err)
	}

	return nil
}

func decode(imageType string, file io.Reader) (image.Image, error) {
	switch imageType {
	case "png":
		return png.Decode(file)
	case "jpeg":
		return jpeg.Decode(file)
	default:
		return image.NewRGBA(image.Rectangle{}), fmt.Errorf("image type '%s' not implemented yet", imageType)
	}
}
