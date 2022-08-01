package chapter7

import "fmt"

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 { //嵌套
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
} //循环

func Chapter7main8() {
	fmt.Println(fib(6))
}
