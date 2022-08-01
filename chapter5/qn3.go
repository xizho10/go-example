package chapter5

import "fmt"

func Chapter5main5() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz") //i mod 3 = 0，输出Fizz
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println("FizzBuzz")
		}
	}
}
