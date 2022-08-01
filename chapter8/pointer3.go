package chapter8

import "fmt"

// use a pointer to change the actual passed in variable
func one(x *int) {
	*x = 1
}

func Chapter8main4() {
	x := new(int)
	one(x)
	fmt.Println(*x)
}
