package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Cpoints.sd")
	if err != nil {
		fmt.Printf("something is teribly teibly wrong with file not created ")
		panic(err)
	}
	var x float64
	for i := 0; i < 786; i++ {
		fmt.Scan(&x)
		binary.Write(file, binary.LittleEndian, &x)
	}

}
