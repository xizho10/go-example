package chapter7

import "fmt"

func half(x int) (int, bool) {
	halved := x / 2

	if x%2 == 0 {
		return halved, true
	} else {
		return halved, false
	}
} //分支

func Chapter7main() {
	fmt.Println(half(1))
}
