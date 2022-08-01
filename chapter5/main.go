package chapter5

import "fmt"

func Chapter5main() {
	i := 1 //i赋值1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	} //在i<=10时，输出i，然后i = i + 1，循环
}
