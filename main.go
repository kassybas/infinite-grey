package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
)

// GimmeImages returns an image
func GimmeImages(offset float64, curTarget float64) {
	iCurTarget := int(curTarget)
	numOfGrey := 2.0

	dirName := "results/" + strconv.FormatFloat(curTarget, 'f', -1, 64)
	os.Mkdir(dirName, os.ModePerm)
	// All the possibilities
	alef := math.Pow(numOfGrey, curTarget*curTarget)
	for i := offset; i < alef; i++ {
		img := image.NewRGBA(image.Rect(0, 0, iCurTarget, iCurTarget))
		for px := 0; px < iCurTarget*iCurTarget; px++ {
			x := px / iCurTarget
			y := px % iCurTarget
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
		fileName := dirName + strconv.FormatFloat(i, 'f', -1, 64) + ".png"

		SaveImage(img, fileName)
	}
}

// SaveImage saves image as PNG
func SaveImage(img image.Image, path string) {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	curTarget := 2.0
	GimmeImages(0, curTarget)
}
