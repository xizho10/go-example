// a go routine is a function
// that is capable of running concurrently with other functions
// use go keyword followed by function invocation

package chapter10

import (
	"fmt"
	"math/rand"
	"time"
)

func k(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(256)) //时间间隔
		time.Sleep(time.Millisecond * amt)
	}
}

func Chapter10main5() {
	// run 10 go routines
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("test: ", input)
}
