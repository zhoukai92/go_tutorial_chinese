package main

import "fmt"

/*
赋值运算
*/
func main() {
	var num1 = 10
	fmt.Println(num1)
	var num2 = (10+20)%3 + 3 - 7 //=右侧的值运算清楚后，再赋值给=的左侧
	fmt.Println(num2)
	var num3 = 10
	num3 += 20 //等价num3 = num3 + 20;
	fmt.Println(num3)
}
