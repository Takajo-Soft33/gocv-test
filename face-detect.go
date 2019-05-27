package main

import (
  "fmt"
  "image"
  "image/color"
  "os"

  "gocv.io/x/gocv"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Usage: go run face-detect.go <image file>\n")
    return
  }

  filename := os.Args[1]
  xmlFile := "opencv/data/haarcascades/haarcascade_frontalface_default.xml"

  img := gocv.IMRead(filename, gocv.IMReadColor)

  window := gocv.NewWindow("Face Detect")
  defer window.Close()

  borderColor := color.RGBA{0, 255, 0, 100}

  classifier := gocv.NewCascadeClassifier()
  defer classifier.Close()

  if ! classifier.Load(xmlFile) {
    fmt.Printf("Error reading cascade file: %v\n", xmlFile)
    return
  }

  size := img.Size()
  rects := classifier.DetectMultiScaleWithParams(img, 1.1, 30, 0, image.Pt(10, 10), image.Pt(size[0], size[1]))
  fmt.Printf("found %d faces\n", len(rects))

  for _, r := range rects {
    gocv.Rectangle(&img, r, borderColor, 3)
    size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
    pt := image.Pt(r.Min.X + (r.Min.X/2) - (size.X/2), r.Min.Y - 2)
    gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, borderColor, 2)
  }

  window.IMShow(img)
  for {
    if window.WaitKey(1) >= 0 {
      break
    }
  }
}
