package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"

	"github.com/kassybas/infinite-grey/src/imgen"

	"github.com/nfnt/resize"
)

// GimmeImages returns an image
func GimmeImages(offset float64, curTarget float64) {
	numOfGrey := 2.0
	dirName := "../results/" + strconv.FormatFloat(curTarget, 'f', -1, 64)
	os.Mkdir(dirName, os.ModePerm)

	// All the possibilities TODO: handle big numbers
	alef := math.Pow(numOfGrey, curTarget*curTarget)
	fmt.Println("You can expect: ", alef)
	for i := offset; i < alef; i++ {
		pixelBitMask := imgen.GetPixelBitMask(curTarget)
		img := imgen.GenerateImage(i, curTarget, pixelBitMask)
		fileName := dirName + "/img_" + strconv.FormatFloat(i, 'f', -1, 64) + ".png"
		SaveImage(img, fileName)
	}
}

// SaveImage resizes and saves image as PNG
func SaveImage(img image.Image, path string) {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	// Docs: https://github.com/nfnt/resize
	m := resize.Resize(480, 480, img, resize.Bilinear)
	png.Encode(f, m)
	fmt.Println("Saved image", path)
}

func main() {
	curTarget := 3.0
	GimmeImages(0, curTarget)
}
