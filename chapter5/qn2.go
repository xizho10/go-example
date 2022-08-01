package chapter5

import "fmt"

func Chapter5main4() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		} //如果i mod 3 = 0，输出i， 然后i++，满足i <= 100，循环
	}
}
