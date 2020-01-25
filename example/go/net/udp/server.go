package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer conn.Close()
	for {
		var data [1024]byte
		n, addr, err := conn.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		_, err = conn.WriteToUDP([]byte(reply), addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
