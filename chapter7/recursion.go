package chapter7

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func Chapter7main13() {
	fmt.Println(factorial(3))
}
