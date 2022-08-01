package chapter10

import (
	"fmt"
	"time"
)

func mysleep(n int) {
	select {
	case <-time.After(time.Duration(n) * time.Second): //返回时间
		fmt.Println("End of Sleep")
	}
}

func chapter10main6() {
	fmt.Println("Start")
	mysleep(3) //延迟
	fmt.Println("End")

	var input string
	fmt.Scanln(&input)
}
