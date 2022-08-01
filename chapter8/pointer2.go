package chapter8

import "fmt"

// use a pointer to change the actual passed in variable
func zero2(x *int) {
	*x = 0
}

func Chapter8main3() {
	x := 5
	zero2(&x)
	fmt.Println(x)
}
