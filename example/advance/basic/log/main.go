package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Log entry")
	log.SetFlags(0)
	for i := 0; i < 100; i++ {
		go log.Println(i)
	}
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	log.SetOutput(ioutil.Discard)
	log.Println("Entry 2")
	defer log.Println("Will not be logged")
	log.Fatal("Exit")
}
