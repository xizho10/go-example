package chapter7

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func Chapter7main7() {
	defer second()
	first()
}
