package main

import "math/rand"

type Point struct {
	x, y, r float64
}

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
