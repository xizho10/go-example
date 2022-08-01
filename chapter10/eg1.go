// a go routine is a function
// that is capable of running concurrently with other functions
// use go keyword followed by function invocation

package chapter10

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	} //循环输出
}

func Chapter10main3() {
	go f(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println("test: ", input)
}
