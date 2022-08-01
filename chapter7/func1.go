package chapter7

import "fmt"

func Chapter7main9() {
	xs := []float64{98, 93, 77, 82, 83} //浮点型，初始化赋值
	total := 0.0

	for _, v := range xs {
		total += v
	}

	fmt.Println(total / float64(len(xs)))
}
