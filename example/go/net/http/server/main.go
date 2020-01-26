package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上(query param)，请求体中是没有数据的
	queryParam := r.URL.Query()    // 自动帮我们识别URL中的参数
	name := queryParam.Get("name") // 可以支持中文
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	// 我在服务端打印客户端发来的body
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/post/Go/15_socket/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
