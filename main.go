package main

import (
	"gocv.io/x/gocv"
)

func rotate() {

}

func main() {

	var CANNY1 float32 = 100.0
	var CANNY2 float32 = 200.0

	img := gocv.IMRead("data/meter1.jpg", gocv.IMReadGrayScale)

	edges := gocv.NewMat()

	gocv.Canny(img, &edges, CANNY1, CANNY2)

	window := gocv.NewWindow("Image")
	defer window.Close()

	window.IMShow(edges)

	for {

		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
