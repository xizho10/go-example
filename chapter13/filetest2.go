package chapter13

import (
	"fmt"
	"io/ioutil"
)

func Chapter13main4() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
