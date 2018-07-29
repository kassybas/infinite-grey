package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

// GimmeImages returns an image
func GimmeImages(offset float64, curTarget float64) {
	iCurTarget := int(curTarget)
	numOfGrey := 2.0

	dirName := "results/" + strconv.FormatFloat(curTarget, 'f', -1, 64)
	os.Mkdir(dirName, os.ModePerm)

	// All the possibilities
	alef := math.Pow(numOfGrey, curTarget*curTarget)
	fmt.Println("You can expect: %f", alef)
	for i := offset; i < alef; i++ {
		img := image.NewRGBA(image.Rect(0, 0, iCurTarget, iCurTarget))

		// Fill up each pixel in the image
		for px := 0; px < iCurTarget*iCurTarget; px++ {

			x := px / iCurTarget
			y := px % iCurTarget

			pixelBit := int(math.Pow(2.0, float64(px)))

			// TODO: fix math
			pxColor := uint8(int(i)&pixelBit) * 255
			// fmt.Println(pxColor)
			img.Set(x, y, color.RGBA{pxColor, pxColor, pxColor, 255})
		}
		fileName := dirName + "/img_" + strconv.FormatFloat(i, 'f', -1, 64) + ".png"

		SaveImage(img, fileName)
	}
}

// SaveImage resizes and saves image as PNG
func SaveImage(img image.Image, path string) {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	// Docs: https: //github.com/nfnt/resize
	m := resize.Resize(480, 480, img, resize.Lanczos3)
	// m := resize.Resize(480, 480, img, resize.NearestNeighbor)
	png.Encode(f, m)
	fmt.Println("Saved image", path)
}

func main() {
	curTarget := 4.0
	GimmeImages(0, curTarget)
}
