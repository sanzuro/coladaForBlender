package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf(`<?xml version="1.0" encoding="utf-8"?>
          	<COLLADA xmlns="http://www.collada.org/2005/11/COLLADASchema" version="1.4.1">
	   
	                     <library_images/>
						 <library_geometries>`)

	for i := 0; i < 4; i++ {
		fmt.Printf(`<geometry id ="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh" name= "Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`" ><mesh>`)
		fmt.Printf(`<source id="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-positions">`)
		fmt.Printf(`<float_array id="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-positions-array" count="12"> `)
		for j := 0; j < 12; j++ {
			fmt.Printf("%d ", rand.Intn(2))
		}
		fmt.Printf(`</float_array>`)
		fmt.Printf(`<technique_common>`)
		fmt.Printf(`<accessor source="#Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-positions-array" count="4" stride="3">`)
		fmt.Printf(`<param name="X" type="float"/>
		<param name="Y" type="float"/>
		<param name="Z" type="float"/>
	        </accessor>
        	</technique_common>
            </source>
			<source id="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-normals">`)
		fmt.Printf(`<float_array id="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-normals-array" count="3">`)
		for j := 0; j < 3; j++ {
			//fmt.Printf("%f ", rand.Float32())
			//fmt.Printf("")
		}
		// this maight be a standerd value
		//subject to change
		fmt.Printf(" 0 0 1 ")
		// -- this
		fmt.Printf(`</float_array>`)
		fmt.Printf(`<technique_common>
					  <accessor source="#Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-normals-array" count="1" stride="3">`)
		fmt.Printf(` <param name="X" type="float"/>
		           <param name="Y" type="float"/>
	            	<param name="Z" type="float"/>
            	     	  </accessor>
              				</technique_common>
						     </source>
								<vertices id="Plane_00`)

		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-vertices">`)
		fmt.Printf(`<input semantic="POSITION" source="#Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh-positions"/>
      		 		 </vertices>
        				<triangles count="2">
        			  <input semantic="VERTEX" source="#Plane-mesh-vertices" offset="0"/>
        			  <input semantic="NORMAL" source="#Plane-mesh-normals" offset="1"/>
						  <p>`)
		for j := 0; j < 12; j++ {
			fmt.Printf("%d ", rand.Intn(2))
		}
		fmt.Printf(`</p>
    	 			   </triangles>
    					  </mesh>
							</geometry>`)

	}
	fmt.Printf(`</library_geometries>`)
	fmt.Printf(`<library_controllers/>`)

	fmt.Printf(` <library_visual_scenes>
					<visual_scene id="Scene" name="Scene">`)
	for i := 0; i < 4; i++ {
		fmt.Printf(`<node id="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`" name="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`" type="NODE">`)

		fmt.Printf(`<instance_geometry url="#Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`-mesh" name="Plane_00`)
		fmt.Printf("%d", i)
		fmt.Printf(`"/>
				      	</node>`)
	}
	fmt.Printf(`  </visual_scene>
	             </library_visual_scenes>
	               <scene>
					  <instance_visual_scene url="#Scene"/>
					</scene>
					  </COLLADA>`)

}
