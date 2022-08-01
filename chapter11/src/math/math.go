package math

import "sort"

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) < 1 {
		return 0
	} //若xs无元素，退出

	total := float64(0)
	for _, x := range xs {
		total += x //循环
	} //分支

	return total / float64(len(xs)) //返回值
}

// Finds the minimum of a series of numbers
func Min(xs []float64) float64 {
	sort.Float64s(xs)
	return xs[0]
} //得到最小值

// Finds the maximum of a series of numbers
func Max(xs []float64) float64 {
	sort.Float64s(xs)
	return xs[len(xs)-1]
} //得到最大值
