package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows a number of methods
// for searching a string
func SearchString() {
	s := "this is a test"
	// returns true because s contains the word this
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.ContainsAny(s, "abc"))
	fmt.Println(strings.HasPrefix(s, "this"))
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"
	// prints [simple string]
	fmt.Println(strings.Split(s, " "))
	// prints "Simple String"
	fmt.Println(strings.Title(s))
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demonstrates how to create
// an io.Reader interface quickly with a string
func StringReader() {
	s := "simple stringn"
	r := strings.NewReader(s)
	// prints s on Stdout
	io.Copy(os.Stdout, r)
}
