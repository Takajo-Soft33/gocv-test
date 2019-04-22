package main

import (
  "fmt"
  "image"
  "image/color"
  "os"

  "gocv.io/x/gocv"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Printf("Usage: go run face-detect.go <image file> <classifier xml file>\n")
    return
  }

  filename := os.Args[1]
  xmlFile := os.Args[2]

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

  rects := classifier.DetectMultiScale(img)
  fmt.Printf("found %d faces\n", len(rects))

  for _, r := range rects {
    gocv.Rectangle(&img, r, borderColor, 3)

    size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
    pt := image.Pt(r.Min.X + (r.Min.X/2) - (size.X/2), r.Min.Y - 2)
    gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, borderColor, 2)
  }

  window.IMShow(img)
  if window.WaitKey(1) >= 0 {
    return
  }
}
