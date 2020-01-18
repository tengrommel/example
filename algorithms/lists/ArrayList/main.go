package main

import (
	"awesomeProject/algorithms/lists/ArrayList/ArrayList"
	"fmt"
)

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a")
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}
