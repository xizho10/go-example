package chapter7

import (
	"fmt"
	"sort"
)

func greatest(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1] //返回值
}

func Chapter7main2() {
	x := []int{300, 3, 2, 1000}
	fmt.Println(greatest(x...))
}
