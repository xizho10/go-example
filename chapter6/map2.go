package chapter6

import "fmt"

func Chapter6main6() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)        //输出所有元素
	fmt.Println(x["key"]) //输出key的那一个元素
}
