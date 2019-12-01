package main

import "fmt"

func main() {
	type XYZ struct {
		X int
		Y int
		Z int
	}
	var s1 XYZ
	fmt.Println(s1.X, s1.Z)
	// As you can see, there is nothing that prevents you from defining a new structure type
	// inside a function, but you should have a reason for doing so.
	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)
	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)
	// when you assign a structure to an array of structures,
	// the structure is copied into the array so changing the value
	// of the original structure will have no effect on the objects of the array
}
