package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// GimmeImage returns an image
func GimmeImage(offset int64) image.Image {
	// Create an 100 x 100 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// Draw a red dot at (2, 3)
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
	return img
}

// SaveImage saves image as PNG
func SaveImage(img image.Image, path string) {
	// Save to out.png
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
