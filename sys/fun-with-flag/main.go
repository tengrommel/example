package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return errors.New("cannot use names flag more than once")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mahalia", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")
	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)
	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}
	fmt.Println("Remaining command line argument:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
