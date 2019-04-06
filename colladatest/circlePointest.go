package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("Cpoints.sd")
	var x float64
	for i := 0; i < 256; i++ {
		fmt.Printf("%v ", i)
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v ", x)
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v ", x)
		binary.Read(file, binary.LittleEndian, &x)
		fmt.Printf("%v ", x)
		fmt.Printf("\n")
	}
}
