package main

import "fmt"

func main() {
	fmt.Printf(`<?xml version="1.0" encoding="utf-8"?>
	<COLLADA xmlns="http://www.collada.org/2005/11/COLLADASchema" version="1.4.1">

	  <library_images/>
	  <library_geometries>`)

	var p, c Point
	var r float64

	Ccircle(p, c, r, Circle265)
}
