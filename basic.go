package main

import (
	"fmt"
	"math"
)

// this is exported
// by many lib

// MeshArray ... this would be the main controlling array
var MeshArray = make([]int, 0)

var (
	Circle265 = int(1)
)

// Point ... the Point or may be vecter
type Point struct {
	X, Y, Z float64
}

// Mod ... this will give mod
func Mod(x Point) float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y + x.Z*x.Z)
}

// Unit ... this is for the unit vecter
func (p Point) Unit() Point {
	var v Point
	t := Mod(p)
	v.X, v.Y, v.Z = p.X/t, p.Y/t, p.Z/t
	return v
}

// Cinit ... this will initialize the collada header

func Cinit() {
	fmt.Printf(` <?xml version="1.0" encoding="utf-8"?>
				 <COLLADA xmlns="http://www.collada.org/2005/11/COLLADASchema" version="1.4.1">
			 	 <library_images/>
	 			 <library_geometries>`)
}

// Cnoder ... this will edit compleate node of collada
func Cnoder() {

	// back content initializer
	fmt.Printf(`    
					</library_geometries>
		            <library_controllers/>
		            <library_visual_scenes>
					<visual_scene id="Scene" name="Scene">
					`)

	for t, x := range MeshArray {

		switch x {
		case Circle265:
			{
				// one compleate node data with indexing of meshes in gieven order

				fmt.Printf(`<node id="Circle_00`)
				fmt.Printf("%d", t)
				fmt.Printf(`" name="Circle_00`)
				fmt.Printf("%d", t)
				fmt.Printf(`" type="NODE">
				<instance_geometry url="#Circle`)
				fmt.Printf("%d", t)
				fmt.Printf(`-mesh" name="Circle`)
				fmt.Printf("%d", t)
				fmt.Printf(`"/>
				</node>`)
				fmt.Printf("\n")
			}
		}
	}

	fmt.Printf(`</visual_scene>
				</library_visual_scenes>
				<scene>
				<instance_visual_scene url="#Scene"/>
				</scene>
				</COLLADA>`)
}
