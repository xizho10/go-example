package chapter13

import (
	"os"
)

func Chapter13main5() {
	file, err := os.Create("./files/zhengtang.txt")
	if err != nil {
		return
	}

	defer file.Close()

	var str string
	str = "Test"
	file.WriteString(str)
}
