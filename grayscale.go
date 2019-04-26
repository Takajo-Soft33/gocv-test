package main

import (
  "os"
  "fmt"

  "gocv.io/x/gocv"
)

func grayscale(filename string) gocv.Mat {
  img := gocv.IMRead(filename, gocv.IMReadColor)
  grayImg := gocv.NewMat()
  gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray)
  return grayImg
}

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Usage: go run grayscale.go <imagefile>\n")
    return
  }
  filename := os.Args[1]
  window := gocv.NewWindow("Grayscale")
  img := grayscale(filename)

  for {
    window.IMShow(img)
    window.WaitKey(1)
  }
}
