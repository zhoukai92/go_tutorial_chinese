package main

import (
	"fmt"
	"unsafe"
)

/*
1. 布尔类型也叫bool类型，bool类型数据只允许取值true和false
2. 布尔类型占1个字节。
3. 布尔类型适于逻辑运算，一般用于程序流程控制
*/
func main() {
	//测试布尔类型的数值：
	var flag01 bool = true
	fmt.Println(flag01, unsafe.Sizeof(flag01))
	var flag02 bool = false
	fmt.Println(flag02)
	var flag03 bool = 5 < 9
	fmt.Println(flag03)
}
