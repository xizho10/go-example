package chapter8

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func Chapter8main() {
	x := 1
	y := 2

	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
