package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("CPoints.sd")
	if err != nil {
		panic(err)
	}
	var x []byte
	file.Read(x)
	fmt.Printf("%+v", x)
}
