package chapter6

import "fmt"

func Chapter6main4() {
	arr := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	x := arr[:5]

	var total float64 = 0

	for _, value := range x {
		total += value
	} //循环

	fmt.Println(total / float64(len(x)))
}
