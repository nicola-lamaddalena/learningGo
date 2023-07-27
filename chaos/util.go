package main

import (
	"math/rand"
	"sync"

	"github.com/fogleman/gg"
)

type Point struct {
	x, y, r float64
}

const (
	width         = 1000
	height        = 1000
	pointRadius   = 1
	steps         = 400000
	numGoroutines = 16 // Number of goroutines to use
)

func midPoint(p1 Point, p2 Point, radius float64) Point {
	return Point{
		(p1.x + p2.x) / 2,
		(p1.y + p2.y) / 2,
		radius,
	}
}

func chooseVertex(vertices [3]Point) Point {
	randIdx := rand.Intn(len(vertices))
	vertex := vertices[randIdx]
	return vertex
}

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

func GoRoutineTriangle() {
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
	dc.SavePNG("./img/sierpinskiTriangleGoroutines.png")
}

func Triangle() {
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
	dc.SavePNG("./img/sierpinskiTriangle.png")
}
