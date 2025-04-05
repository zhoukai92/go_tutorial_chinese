package main

// 如果有多个包，建议一次性导入,格式如下：
import (
	format "fmt"
	"unsafe"
)

/*
包说明

	函数名，变量名首字母大写，函数，变量可以被其它包访问
	一个目录下不能有重复的函数
	包名和文件夹的名字，可以不一样
	一个目录下的同级文件归属一个包, 同级别的源文件的包的声明必须一致
*/
func main() {
	format.Println("可以给包取别名，取别名后，原来的包名就不能使用了")
	//fmt.Println("error")
	a := 1
	format.Println(unsafe.Sizeof(a))
}

//一个目录下不能有重复的函数
// 'mySum' redeclared in this package
//func mySum(n1 int, n2 int) int {
//	return n1 + n2
//}
