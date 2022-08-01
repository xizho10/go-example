package chapter8

import "fmt"

func zero(x int) {
	x = 0
}

func Chapter8main2() {
	x := 5
	zero(x)
	fmt.Println(x)
}
