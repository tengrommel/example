package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type AngryReader struct {
	r io.Reader
}

func (a *AngryReader) Read(b []byte) (int, error) {
	n, err := a.r.Read(b)
	for r, i, w := rune(0), 0, 0; i < n; i += w {
		r, w = utf8.DecodeRune(b[i:])
		if !unicode.IsLetter(r) {
			continue
		}
		ru := unicode.ToUpper(r)
		if wu := utf8.EncodeRune(b[i:], ru); w != wu {
			return n, fmt.Errorf("%c->%c, size mismatch %d->%d", r, ru, w, wu)
		}
	}
	return n, err
}

func NewAngryReader(r io.Reader) *AngryReader {
	return &AngryReader{r: r}
}

/**
We want to implement a custom reader that takes the content form another reader and
transforms it into uppercase;
*/
func main() {
	a := NewAngryReader(strings.NewReader("Hello, playground! 牛 🐂"))
	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
}
