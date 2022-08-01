// channels allow two go routines to communicate
// with each other and synchronize their actions

package chapter10

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	} //给c赋值
}

func printer(c <-chan string) {
	for {
		msg := <-c //读取c的内容
		fmt.Println("PRINTER: ", msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Chapter10main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input) //输入
}
