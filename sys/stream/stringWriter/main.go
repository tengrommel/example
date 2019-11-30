package main

import (
	"fmt"
	"strings"
)

func main() {
	b := strings.Builder{}
	b.WriteString("one")
	c := b
	b.WriteString("Hey!") // panic: strings: illegal use of one-zero Builder copied by value
	fmt.Println(c.String())
	fmt.Println(b.String())
}
