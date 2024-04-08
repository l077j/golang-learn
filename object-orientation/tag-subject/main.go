package main

import (
	"encoding/json"
	"fmt"
)

// 如果没有 omitempty 标记，则对于不包含 Age 的 JSON 对象，将填充默认值 0，而对于空字符串位置项，也将序列化成空字符串。
// 如果使用 omitempty，Go 编码器将忽略 Person 结构中未设置的字段
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1没有Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// =======================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2会打印出来Addr
	fmt.Printf("%s\n", data2)
}
