package chapter6

import "fmt"

func Chapter6main2() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) //设定长度为2，为整形
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
