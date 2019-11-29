package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please specify a source and a dest")
		return
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer src.Close()
	// OpenFile allows to open a file with any permissions
	dst, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dst.Close()
	cur, err := src.Seek(0, io.SeekEnd)
	// Let's go to the end of the file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b := make([]byte, 16)
	for step, r, w := int64(16), 0, 0; cur != 0; {
		if cur < step { // ensure cursor is 0 at max
			b, step = b[:cur], cur
		}
		// 确定相对位移
		cur = cur - step
		// 将指针指向开始点
		_, err = src.Seek(cur, io.SeekStart) // go backwards
		if err != nil {
			break
		}
		// 读取文件内容进入b byte数组
		if r, err = src.Read(b); err != nil || r != len(b) {
			if err == nil {
				err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
			}
			break
		}
		// 交换反转
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j+1 {
			switch {
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}
			b[i], b[j] = b[j], b[i] // swap bytes
		}
		// 写入文件
		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}
	if err != nil && err != io.EOF { // we expect an EOF
		fmt.Println("\n\nError:", err)
	}
}
