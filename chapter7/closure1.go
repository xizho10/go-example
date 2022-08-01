package chapter7

import "fmt"

func Chapter7main4() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))
}
