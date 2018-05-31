package main

import (
	"gocv.io/x/gocv"
)

func process(filePath string) gocv.Mat {
	var CANNY1 float32 = 100.0
	var CANNY2 float32 = 200.0

	img := gocv.IMRead("data/meter1.jpg", gocv.IMReadReducedGrayscale4)

	edges := gocv.NewMat()

	gocv.Canny(img, &edges, CANNY1, CANNY2)

	return edges
}

func main() {

	window := gocv.NewWindow("Image")
	defer window.Close()

	image := process("data/meter1.jpg")

	window.IMShow(image)

	for {

		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
