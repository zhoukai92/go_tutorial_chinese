package main

import "fmt"

/*
变量的使用步骤：
1.声明
2.赋值
3.使用
*/
func main() {
	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age = ", age)
	//声明和赋值可以合成一句：
	var age2 int = 19
	fmt.Println("age2 = ", age2)
	// 不可以在赋值的时候给与不匹配的类型
	//var num int = 12.56
	//fmt.Println("num = ", num)
}
