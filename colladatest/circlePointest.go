package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// this file has only has 399 values
// tested

func main() {
	file, err := os.Open("CPoints.sd")
	if err != nil {
		panic(err)
	}
	var x float64
	for i := 0; i < 400; i++ {
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v ", x)
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v ", x)
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v \n", x)
	}
}
