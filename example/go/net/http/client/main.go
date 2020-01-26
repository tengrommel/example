package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http Client
var client = http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       0,
}

func main() {
	//resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=中文&age=18")
	//if err != nil {
	//	fmt.Printf("get url failed, err: %v\n", err)
	//	return
	//}
	//从resp中把服务端返回的数据读出来
	//b, err := ioutil.ReadAll(resp.Body) // 我在客户端打印服务端相应的body
	//if err != nil {
	//	fmt.Printf("read resp.Body failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))
	data := url.Values{} // url values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周林")
	data.Set("age", "9000")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed, err: %v\n", err)
	//	return
	//}
	// 禁用KeepAlive的client

	resp, err := client.Do(req)
	defer resp.Body.Close()             // 一定要记着关闭
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端打印服务端相应的body
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(b)
}
