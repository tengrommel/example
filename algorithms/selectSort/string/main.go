package main

import "fmt"

func main() {
	fmt.Println("a" > "b")
	fmt.Println("a" < "b")
	fmt.Println("a1" == "a")
	pa := "a"
	pb := "b"
	fmt.Println(&pa, &pb)
	fmt.Println(pa < pb)
}
