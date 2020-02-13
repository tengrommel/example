package main

import (
	"awesomeProject/cookbook/flags"
	"flag"
	"fmt"
)

func main() {
	c := flags.Config{}
	c.Setup()
	flag.Parse()
	fmt.Println(c.GetMessage())
}
