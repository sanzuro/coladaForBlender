package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Point ...  the generated file would have 399 points
type Point struct {
	X, Y, Z float64
}

func main() {
	file, err := os.Open("CPoints.sd")
	// for creating file
	// replace Open with Create
	if err != nil {
		fmt.Printf("sex happened")
		panic(err)
	}
	defer file.Close()
	i := 0
	for x := float64(-1.0000); x < 1; x += float64(0.01) {
		i++
		var n Point
		// n.X = x
		// n.Z = 0
		// n.Y = math.Sqrt(1 - x*x)
		// binary.Write(file, binary.LittleEndian, &n)
		// fmt.Printf(" %v %9.2v %v %v \n", i, n.X, n.Y, n.Z)
		binary.Read(file, binary.LittleEndian, &n)
		fmt.Printf("%v %v %v %v \n", i, n.X, n.Y, n.Z)
	}
	for x := float64(0.99); x > -1; x -= float64(0.01) {
		var n Point
		i++
		// n.X = x
		// n.Z = 0
		// n.Y = -math.Sqrt(1 - x*x)
		// binary.Write(file, binary.LittleEndian, &n)
		// fmt.Printf("   %v %9.2v %v %v \n", i, n.X, n.Y, n.Z)
		binary.Read(file, binary.LittleEndian, &n)
		fmt.Printf("%v %v %v %v \n", i, n.X, n.Y, n.Z)
	}

}
