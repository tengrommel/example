package main

import (
	"awesomeProject/algorithms/lists/ArrayList/ArrayList"
	"fmt"
)

func fab(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fab(num-1) + fab(num-2)
	}
}

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("b3")
	list.Append("d3")
	list.Append("f3")
	list.Append("e3")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "d3" {
			it.Remove()
		}
		fmt.Println(item)
	}
}
