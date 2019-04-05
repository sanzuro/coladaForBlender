package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

// using the trick of standerd values and the
// rotating and then roatating them
// via rotaional matricies
// here the initvalues is the default
// value for Points , circle at centre and facing upwards
var initValues [400]Point

var xCosAngle, yCosAngle, zCosAngle float64
var xSinAngle, ySinAngle, zSinAngle float64

// this is general roatation matrix for x angle
// general implimentation
func xRot(x Point) Point {
	var v Point
	v.X = x.X
	v.Y = xCosAngle*x.Y + xSinAngle*x.Z
	v.Z = xCosAngle*x.Z - xSinAngle*x.Y
	//fmt.Printf("%+v\n", v)
	return v
}

//same way implimented y rotation matrix
// general implimentation
func yRot(y Point) Point {
	var v Point
	v.Y = y.Y
	v.X = yCosAngle*y.X - ySinAngle*y.Z
	v.Z = ySinAngle*y.X + yCosAngle*y.Z
	//fmt.Printf("%+v\n", v)
	return v
}

// zRot func implimented in the same way
func zRot(z Point) Point {
	var v Point
	v.Z = z.Y
	v.X = zCosAngle*z.X - zSinAngle*z.Y
	v.Y = zCosAngle*z.X + zSinAngle*z.Y
	//fmt.Printf("%+v\n", v)
	return v
}

func init() {
	file, err := os.Open("CPoints.sd")
	if err != nil {
		fmt.Printf("file not found : CPoints.sd\n")
		panic(err)
	}
	for i := 1; i < 400; i++ {
		binary.Read(file, binary.LittleEndian, &initValues[i])
		// uncomment this following line for test proper loading
		// fmt.Printf("%v : %v\n", i+1, initValues[i])
	}
}

// Ccircle ...  this is the real sh**
// the proper formatted circle geomatry argument
func Ccircle(c Point, p Point, r float64) {
	mod := Mod(p)
	xCosAngle, yCosAngle, zCosAngle = p.X/mod, p.Y/mod, p.Z/mod
	xSinAngle, ySinAngle, zSinAngle = math.Sqrt(1-xCosAngle*xCosAngle), math.Sqrt(1-yCosAngle*yCosAngle), math.Sqrt(1-zCosAngle*zCosAngle)

	for i, iVal := range initValues[:] {
		_ = i
		_ = iVal

		n := xRot(yRot(zRot(iVal)))
		fmt.Printf(" %v : %+v\n", i, n)
		// to rotate every value in the
		// slice with gievn angle of p (vecter)
	}
}

func main() {
	t := Point{2.3, 6.7, 8.9}
	z := Point{1.8, 6.7, 9.0}
	fmt.Printf("%+v %+v\n", t, z)
	Ccircle(t, z, 9.0008)
}
