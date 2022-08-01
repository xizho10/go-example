package chapter7

import "fmt"

func Chapter7main5() {
	x := 0

	increment := func() int {
		x++ //x+1
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
