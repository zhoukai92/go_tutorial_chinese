package main

import "fmt"

/*
指针运算
*/
func main() { //定义一个变量：
	var age = 22
	fmt.Println("age对应的存储空间的地址为：", &age) //age对应的存储空间的地址为： 0xc0000100b0
	var ptr = &age
	fmt.Println(ptr, ptr == &age)
	fmt.Println("ptr这个指针指向的具体数值为：", *ptr)

}
