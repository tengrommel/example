package main

import (
	"awesomeProject/cookbook/interfaces"
	"bytes"
	"fmt"
)

// The Copy() function copies bytes between interfaces and treats that data like a stream.
// Thinking of data as streams has a lot of practical uses, especially when working with
// network traffic or filesystems. The Copy() function also creates a MultiWriter
// interface that combines two writer streams and writes
// to them twice using ReadSeeker.

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
