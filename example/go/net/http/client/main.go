package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http Client

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=中文&age=18")
	if err != nil {
		fmt.Printf("get url failed, err: %v\n", err)
		return
	}
	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端打印服务端相应的body
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
