package main

import "math"

// this is exported
// by many lib

type Point struct {
	X, Y, Z float64
}

func Mod(x Point) float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y + x.Z*x.Z)
}
