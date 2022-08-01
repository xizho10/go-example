package chapter5

import "fmt"

func Chapter5main3() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // gets the remainder//如果i mod 2等于0输出even，否则输出odd
			fmt.Println(i, "even") //输出偶数
		} else {
			fmt.Println(i, "odd") //输出奇数
		} //结束后i+1，如果i满足i <= 10，循环
	}
}
