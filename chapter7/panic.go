package chapter7

import "fmt"

func Chapter7main12() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}
