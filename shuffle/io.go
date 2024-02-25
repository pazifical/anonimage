package shuffle

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func loadImage1D(imagePath string) (Image1D, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return Image1D{}, err
	}
	defer f.Close()

	_, imageType, err := image.Decode(f)
	if err != nil {
		return Image1D{}, err
	}

	f.Seek(0, 0)
	loadedImage, err := decode(imageType, f)
	if err != nil {
		return Image1D{}, err
	}

	width := loadedImage.Bounds().Max.X
	height := loadedImage.Bounds().Max.Y

	pixels := make([]color.Color, width*height)

	for y := range height {
		for x := range width {
			pixels[y*width+x] = loadedImage.At(x, y)
		}
	}
	return Image1D{
		Pixels: pixels,
		Width:  width,
		Height: height,
	}, nil
}

func writeImage1D(image1d Image1D, outputPath string) error {
	tempImage := image.NewRGBA(image.Rect(0, 0, image1d.Width, image1d.Height))

	for y := range image1d.Height {
		for x := range image1d.Width {
			tempImage.Set(x, y, image1d.Pixels[y*image1d.Width+x])
		}
	}

	writer, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer writer.Close()

	err = png.Encode(writer, tempImage)
	if err != nil {
		writer.Close()
		return err
	}

	return nil
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
