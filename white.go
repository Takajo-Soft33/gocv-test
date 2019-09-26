package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

// WhiteToBG replaces white area of fg (.jpg) with bg (.jpg)
func WhiteToBG(fg, bg string) {
	fgImg, err := imgio.Open("fg.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	rct := fgImg.Bounds()
	width := rct.Max.X
	height := rct.Max.Y
	rect := image.Rect(0, 0, width, height)

	newImg := image.NewRGBA(rect)

	// low saturation, high value -> transparent
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := fgImg.At(x, y).RGBA()
			r8 := uint8(r / 257)
			g8 := uint8(g / 257)
			b8 := uint8(b / 257)
			a8 := uint8(a / 257)
			max := max(r8, max(g8, b8))
			min := min(r8, min(g8, b8))
			if r8 > 150 && g8 > 150 && b8 > 150 && float64(max-min)/float64(max) < 0.2 {
				a8 = 150
			}
			newImg.Set(x, y, color.RGBA{r8, g8, b8, a8})
		}
	}

	bgImg, err := imgio.Open("bg.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bgImg = transform.Resize(bgImg, width, height, transform.Linear)

	result := blend.Normal(bgImg, newImg)

	if err := imgio.Save("output.jpg", result, imgio.JPEGEncoder(90)); err != nil {
		fmt.Println(err.Error())
		return
	}
}

type changeable interface {
	Set(x, y int, c color.Color)
}

func max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
