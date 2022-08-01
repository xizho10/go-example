package chapter4

import "fmt"

func Chapter4main() { //首字母要大写，外部才能调用
	fmt.Print("Enter a number:")
	var input float64       //浮点型
	fmt.Scanf("%f", &input) //scan 意思就是等待我们输入命令，再点击enter。scan是个单词，f是另一个意思

	output := input * 2 //输入的值乘2

	fmt.Println(output)
}
