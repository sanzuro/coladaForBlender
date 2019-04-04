package main

import (
	"fmt"
	"math"
)

// using the trick of standerd values and the
// rotating and then roatating them
// via rotaional matricies
// here the initvalues is the default
// value for Points , circle at centre and facing upwards
var initValues []Point

var xCosAngle, yCosAngle, zCosAngle float64
var xSinAngle, ySinAngle, zSinAngle float64

// this is general roatation matrix for x angle
// general implimentation
func xRot(x Point) Point {
	var v Point
	v.Y = xCosAngle*x.Y + xSinAngle*x.Z
	v.Z = xCosAngle*x.Z - xSinAngle*x.Y
	return v
}

//same way implimented y rotation matrix
// general implimentation
func yRot(y Point) Point {
	var v Point
	v.X = yCosAngle*y.X - ySinAngle*y.Z
	v.Z = ySinAngle*y.X + yCosAngle*y.Z
	return v
}

// zRot func

func init() {

}

//this function would do all the things to print
// the proper formatted circle geomatry argument

func Ccircle(c Point, p Point, r float64) {
	mod := Mod(p)
	xCosAngle, yCosAngle, zCosAngle = p.X/mod, p.Y/mod, p.Z/mod
	xSinAngle, ySinAngle, zSinAngle = math.Sqrt(1-xCosAngle*xCosAngle), math.Sqrt(1-yCosAngle*yCosAngle), math.Sqrt(1-zCosAngle*zCosAngle)
	for i, iVal := range initValues {
		_ = i
		_ = iVal
		// to rotate every value in the
		// slice with gievn angle of p (vecter)
	}
}

func main() {
	fmt.Print("hello world  \n")
	t := Point{2.3, 6.7, 8.9}
	fmt.Printf("%+v", t)
}
