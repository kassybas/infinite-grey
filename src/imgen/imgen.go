package imgen

import (
	"image"
	"image/color"
	"math"
)

const (
	darkPixel = 255
)

func GetPixelBitMask(size float64) []float64 {
	isize := uint64(size)
	pixelBitMask := make([]float64, isize*isize)
	for i := uint64(0); i < isize*isize; i++ {
		pixelBitMask[i] = math.Pow(2.0, float64(i))
	}
	return pixelBitMask
}

// GenerateImage Generates an image based
func GenerateImage(seed float64, size float64, pixelBitMask []float64) image.Image {
	isize := int(size)
	img := image.NewRGBA(image.Rect(0, 0, isize, isize))

	pixelBase := uint64(seed + 1)
	// Fill up each pixel in the image
	for px := 0; px < isize*isize; px++ {
		x := px / isize
		y := px % isize

		// Get this permutations value for each pixel (divide with that bit's positional value to normalizeit to 1)
		pixelBitVal := uint64(pixelBitMask[px])
		isDark := pixelBase & pixelBitVal / pixelBitVal

		col := uint8(isDark * darkPixel)
		img.Set(x, y, color.RGBA{col, col, col, 255})
	}
	return img
}
