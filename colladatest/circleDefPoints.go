package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Point struct {
	X, Y, Z float64
}

func main() {
	file, err := os.Open("CPoints.sd")
	if err != nil {
		fmt.Printf("sex happened")
		panic(err)
	}
	defer file.Close()
	for x := float64(-1.0000); x < 1; x += float64(0.01) {
		var n Point
		// n.x = x
		// n.z = 0
		// n.y = math.Sqrt(1 - x*x)
		// binary.Write(file, binary.LittleEndian, &n)
		binary.Read(file, binary.LittleEndian, &n)
		fmt.Printf("%v %v %v \n", n.X, n.Y, n.Z)
	}

}
