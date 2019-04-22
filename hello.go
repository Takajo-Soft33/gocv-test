package main

import (
  "os"

  "gocv.io/x/gocv"
)

func main() {
  //webcam, _ := gocv.VideoCaptureDevice(0)
  filename := os.Args[1]
  window := gocv.NewWindow("Hello")
  img := gocv.IMRead(filename, gocv.IMReadColor)

  for {
    //webcam.Read(&img)
    window.IMShow(img)
    window.WaitKey(1)
  }
}
