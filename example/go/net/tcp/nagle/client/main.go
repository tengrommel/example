package main

import (
	"awesomeProject/doc/tcp/proto"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		// 调用协议
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode failed, err:", err)
			return
		}
		conn.Write([]byte(b))
	}
}
