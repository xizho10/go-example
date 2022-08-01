// a go routine is a function
// that is capable of running concurrently with other functions
// use go keyword followed by function invocation

package chapter10

import "fmt"

func g(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func Chapter10main4() {
	// run 10 go routines
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("test: ", input)
}
