package main

import "fmt"

/*
goto: 调整到指定的标签处
不建议使用goto语句
*/
func main() {
	fmt.Println(1)
	fmt.Println(2)
	if true {
		goto label1 //goto一般配合条件结构一起使用
	}
	fmt.Println(3)
label1:
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)

}
