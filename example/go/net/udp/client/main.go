package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败, err:", err)
		return
	}
	defer conn.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		conn.Write([]byte(msg))
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("recv reply msg failed, err:", err)
			return
		}
		fmt.Println("收到回复信息：", string(reply[:n]))
	}
}
