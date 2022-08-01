package chapter7

import "fmt"

func average(xs []float64) float64 {
	// panic("Not implemented!")
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs)) //返回total / float64(len(xs))的值
}

func Chapter7main10() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
