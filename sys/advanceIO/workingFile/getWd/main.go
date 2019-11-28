package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Getting and setting the working directory

/**
We can use the func Getwd() (dir string, err error) function of the os package to find out which path
represents the current working directory.

Changing the working directory is done with another function of the same package,
that is, func Chdir(dir string) error, as shown in the following code:
*/
func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("starting dir:", wd)
	if err := os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}
	if wd, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)
	fmt.Println(filepath.Abs("/Users"))
}
