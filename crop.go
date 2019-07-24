package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sort"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run crop.go <image file>\n")
		return
	}

	filename := os.Args[1]

	img := gocv.IMRead(filename, gocv.IMReadColor)
	gocv.IMWrite("cv.jpg", img)

	gray := gocv.NewMat()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	const t1, t2 float32 = 10, 100

	edges := gocv.NewMat()
	gocv.Canny(gray, &edges, t1, t2)
	gocv.Dilate(edges, &edges, gocv.NewMat())
	gocv.Erode(edges, &edges, gocv.NewMat())

	type contourInfo struct {
		contour []image.Point
		area    float64
	}
	const RetrList = 1
	const ChainApproxNone = 1
	contours := gocv.FindContours(edges, RetrList, ChainApproxNone)

	contourInfos := make([]contourInfo, 0, 10)

	for _, contour := range contours {
		contourInfos = append(contourInfos, contourInfo{contour, gocv.ContourArea(contour)})
	}

	sort.Slice(contourInfos, func(i, j int) bool {
		return contourInfos[i].area > contourInfos[j].area
	})

	//maxContour := contourInfos[0]

	//convex := gocv.NewMat()
	//gocv.ConvexHull(maxContour.contour, &convex, true, true)
	monochrome := gocv.NewMatFromScalar(gocv.Scalar{0, 0, 0, 0}, gocv.MatTypeCV8U)
	gocv.FillPoly(&monochrome, contours, color.RGBA{255, 255, 255, 255})

	window := gocv.NewWindow("Crop")
	defer window.Close()
	window.IMShow(edges)
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
