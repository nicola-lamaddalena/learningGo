package main

import (
	"math/rand"
	"sync"

	"github.com/fogleman/gg"
)

func generatePoints(vertices [3]Point, steps int, pointsChan chan<- Point, wg *sync.WaitGroup) {
	defer wg.Done()

	firstRandomPoint := Point{
		rand.Float64() * width,
		rand.Float64() * height,
		10,
	}
	for i := 0; i < steps; i++ {
		v := chooseVertex(vertices)
		newPoint := midPoint(firstRandomPoint, v, pointRadius)
		pointsChan <- newPoint
		firstRandomPoint = newPoint
	}
}

const (
	width         = 1000
	height        = 1000
	pointRadius   = 1
	steps         = 400000
	numGoroutines = 16 // Number of goroutines to use
)

func main() {
	var vertices = [3]Point{
		{100, 900, pointRadius},
		{900, 900, pointRadius},
		{500, 100, pointRadius},
	}

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)

	pointsChan := make(chan Point, steps)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go generatePoints(vertices, steps/numGoroutines, pointsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(pointsChan)
	}()

	points := []Point{}
	for p := range pointsChan {
		points = append(points, p)
	}

	for i := range points {
		dc.DrawPoint(points[i].x, points[i].y, points[i].r)
	}

	dc.Fill()
	dc.SavePNG("sierpinskiTriangleGoroutines.png")
}
