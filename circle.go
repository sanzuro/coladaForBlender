package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

//the new version is up and running but still has downsides
//fix the --angle thing in the main section

// using the trick of standerd values and the
// rotating and then roatating them
// via rotaional matricies
// here the initvalues is the default
// value for Points , circle at centre and facing upwards

// this mode is made for only 256 point circle
// on later versions the other mode might be added

var initValues [256]Point

var rotAngle float64
var unit Point

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
	file, err := os.Open("Cpoints.sd")
	if err != nil {
		fmt.Printf("file not found : CPoints.sd\n")
		panic(err)
	}
	for i := 0; i < 256; i++ {
		binary.Read(file, binary.LittleEndian, &initValues[i])
		// uncomment this following line for test proper loading
		// fmt.Printf("%v : %v\n", i+1, initValues[i])
	}
}

// Ccircle ...  this is the real circle calling function
// the proper formatted circle geomatry argument
func Ccircle(c Point, p Point, r float64, std int) {

	// obsolete method
	// might be used for some special cases
	// mod := Mod(p)
	// xCosAngle, yCosAngle, zCosAngle = math.Cos(p.X/mod), math.Cos(p.Y/mod), math.Cos(p.Z/mod)
	// xSinAngle, ySinAngle, zSinAngle = math.Sin(math.Sqrt(1-xCosAngle*xCosAngle)), math.Sin(math.Sqrt(1-yCosAngle*yCosAngle)), math.Sin(math.Sqrt(1-zCosAngle*zCosAngle))

	index := len(MeshArray)

	fmt.Printf("\n")
	fmt.Printf(`<geometry id="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh" name="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`">
							<mesh>
							<source id="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-positions">
							<float_array id="Circle_00`)
	fmt.Printf("%d", index)
	//fmt.Printf(`-mesh-positions-array" count="768"> `)
	fmt.Printf(`-mesh-positions-array" count="`)
	switch std {
	case Circle265:
		{
			fmt.Printf("786")
			break
		}
	}
	fmt.Printf(`"> `)

	for i, iVal := range initValues[:] {
		// to rotate every value in the
		// slice with gievn angle of p (vecter)
		_ = i
		_ = iVal
		//fmt.Printf("%+v \n", iVal)
		// n := xRot(yRot(zRot(iVal)))
		n := Rot(iVal)
		fmt.Printf("  %f %f %f", n.X, n.Y, n.Z)

	}

	// after the  float array is done
	// fmt.Printf(`	<technique_common>
	// 							<accessor source="#Circle-mesh-positions-array" count="256" stride="3">
	// 							<param name="X" type="float"/>
	// 							<param name="Y" type="float"/>
	// 							<param name="Z" type="float"/>
	// 							</accessor>
	// 							</technique_common>
	// 							</source>
	// 							<source id="Circle_00`)
	fmt.Printf(`<technique_common>
              <accessor source="#Circle-mesh-positions-array" count="`)

	switch std {
	case Circle265:
		{
			fmt.Printf("256")
			break
		}
	}
	fmt.Printf(`"stride="3">
              <param name="X" type="float"/>
              <param name="Y" type="float"/>
              <param name="Z" type="float"/>
              </accessor>
              </technique_common>
              </source>
              <source id="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-normals">
							<float_array id="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-normals-array" count="0"/>
							<technique_common>
							<accessor source="#Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-normals-array" count="0" stride="3">
							<param name="X" type="float"/>
              <param name="Y" type="float"/>
	            <param name="Z" type="float"/>
            	</accessor>
              </technique_common>
              </source>
              <vertices id="Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-vertices">
              <input semantic="POSITION" source="#Circle`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-positions"/>
              </vertices>
              <lines count="`)
	switch std {
	case Circle265:
		{
			fmt.Printf("256")
			break
		}
	}
	fmt.Printf(`">
             <input semantic="VERTEX" source="#Circle_00`)
	fmt.Printf("%d", index)
	fmt.Printf(`-mesh-vertices" offset="0"/>
             <p>`)

	Cliner(std)
	fmt.Printf(`</p>
              </lines>
              </mesh>
              </geometry>`)

}

// Rot ... this is combination of indevidual roations
func Rot(t Point) Point {
	var v Point
	ca := math.Cos(rotAngle)
	sa := math.Sin(rotAngle)
	var x, y, z float64
	x, y, z = unit.X, unit.Y, unit.Z

	// stuff
	// the new rotation technique based on quaternion formulas and other things
	// cirulation the points over this

	v.X = (ca+x*x*(1-ca))*t.X + (x*y*(1-ca)-z*sa)*t.Y + (x*z*(1-ca)-y*sa)*t.Z
	v.Y = (y*x*(1-ca)+z*sa)*t.X + (ca+y*y*(1-ca))*t.Y + (y*z*(1-ca)-x*sa)*t.Z
	v.Z = (z*x*(1-ca)+y*sa)*t.X + (z*y*(1-ca)-x*sa)*t.Y + (ca+z*z*(1-ca))*t.Z
	return v
}

// Cliner ... this would be used to put the end data in the collada final file
func Cliner(i int) {
	switch i {
	case Circle265:
		{
			fmt.Printf(`0 128 1 129 2 130 3 131 4 132 5 133 6 134 7 135 8 136 9 137 10 138 11 139 12 140 13 141 14 142 15 143 16 144 17 145 18 146 19 147 20 148 21 149 22 150 23 151 24 152 25 153 26 154 27 155 28 156 29 157 30 158 31 159 32 160 33 161 34 162 35 163 36 164 37 165 38 166 39 167 40 168 41 169 42 170 43 171 44 172 45 173 46 174 47 175 48 176 49 177 50 178 51 179 52 180 53 181 54 182 55 183 56 184 57 185 58 186 59 187 60 188 61 189 62 190 63 191 64 192 65 193 66 194 67 195 68 196 69 197 70 198 71 199 72 200 73 201 74 202 75 203 76 204 77 205 78 206 79 207 80 208 81 209 82 210 83 211 84 212 85 213 86 214 87 215 88 216 89 217 90 218 91 219 92 220 93 221 94 222 95 223 96 224 97 225 98 226 99 227 100 228 101 229 102 230 103 231 104 232 105 233 106 234 107 235 108 236 109 237 110 238 111 239 112 240 113 241 114 242 115 243 116 244 117 245 118 246 119 247 120 248 121 249 122 250 123 251 124 252 125 253 126 254 127 255 128 64 129 65 130 66 131 67 132 68 133 69 134 70 135 71 136 72 137 73 138 74 139 75 140 76 141 77 142 78 143 79 144 80 145 81 146 82 147 83 148 84 149 85 150 86 151 87 152 88 153 89 154 90 155 91 156 92 157 93 158 94 159 95 160 96 161 97 162 98 163 99 164 100 165 101 166 102 167 103 168 104 169 105 170 106 171 107 172 108 173 109 174 110 175 111 176 112 177 113 178 114 179 115 180 116 181 117 182 118 183 119 184 120 185 121 186 122 187 123 188 124 189 125 190 126 191 127 192 32 193 33 194 34 195 35 196 36 197 37 198 38 199 39 200 40 201 41 202 42 203 43 204 44 205 45 206 46 207 47 208 48 209 49 210 50 211 51 212 52 213 53 214 54 215 55 216 56 217 57 218 58 219 59 220 60 221 61 222 62 223 63 224 1 225 2 226 3 227 4 228 5 229 6 230 7 231 8 232 9 233 10 234 11 235 12 236 13 237 14 238 15 239 16 240 17 241 18 242 19 243 20 244 21 245 22 246 23 247 24 248 25 249 26 250 27 251 28 252 29 253 30 254 31 255 0</p>
			`)
			return
		}
	default:
		return
	}
}
