package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

/**
snowflake算法：
首先确定数值是64位，int64类型，被划分为4部分，不含开头的第一位，因为这个位是符号位。
随后用41位来表示收到请求时的时间戳，单位为毫秒，然后用5位来表示数据中心的ID，再用5位来表示机器的实例ID，
最后是12位的循环自增ID（到达1111 1111 1111后会归零）。
*/

func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
			"\n",
		)
	}
}
