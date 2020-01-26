package main

import (
	"flag"
	"fmt"
	"time"
)

// flag获取命令行参数
func main() {
	// 创建一个标志位参数
	var name string
	flag.StringVar(&name, "name", "网页", "请输入名字")
	age := flag.Int("age", 90, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")
	// 使用flag
	flag.Parse()
	fmt.Println(name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
}
