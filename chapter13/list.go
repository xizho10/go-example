package chapter13

import (
	"container/list"
	"fmt"
)

func Chapter13main6() {
	// var x list.List
	x := list.New()
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
