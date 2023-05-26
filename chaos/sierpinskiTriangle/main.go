package main

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type Point struct {
	x float64
	y float64
	r float64
}

func main() {
	const width = 1000
	const height = 1000
	const pointRadius = 1
	const steps = 100000

	var vertices = [3]Point{
		{100, 900, 10},
		{900, 900, 10},
		{500, 100, 10},
	}

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)

	firstRandomPoint := Point{
		rand.Float64() * width,
		rand.Float64() * height,
		10,
	}
	points := []Point{}
	for i := 0; i < steps; i++ {
		v := chooseVertex(vertices)
		newPoint := midPoint(firstRandomPoint, v, pointRadius)
		points = append(points, newPoint)
		firstRandomPoint = newPoint
	}
	for i := range points {
		dc.DrawPoint(points[i].x, points[i].y, points[i].r)
	}

	dc.Fill()
	dc.SavePNG("sierpinskiTriangle.png")
}
