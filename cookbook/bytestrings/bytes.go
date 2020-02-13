package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer will make use of the buffer created by the Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)
	// we can quickly convert a buffer back into bytes with b.Bytes() or a
	// string with b.String()
	fmt.Println(b.String())
	// because this is an io Reader we can make use of generic io reader functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	// we can also take our bytes and create a bytes reader
	// these readers implements io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner
	reader := bytes.NewReader([]byte(rawString))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
