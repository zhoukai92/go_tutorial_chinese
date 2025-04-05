package main

import (
	"fmt"
)

/*
字符串是一串固定长度的字符连接起来的字符序列。
*/
func main() {
	//1.定义一个字符串：
	var s1 string = "你好全面拥抱Golang"
	fmt.Println(s1)
	//2.字符串是不可变的：指的是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	s2 = "def"
	// 报错：cannot assign to s2[0] (neither addressable nor a map index expression)
	//s2[0] = 't'
	fmt.Println(s2)

	//3.字符串的表示形式：
	//（1）如果字符串中没有特殊字符，字符串的表示形式用双引号
	//var s3 string = "asdfasdfasdf"
	//（2）如果字符串中有特殊字符，字符串的表示形式用反引号 ``
	var s4 string = `package main

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
`
	fmt.Println(s4)
	//4.字符串的拼接效果：
	var s5 string = "abc" + "def"
	s5 += "hijk"
	fmt.Println(s5)
	//当一个字符串过长的时候：注意：+保留在上一行的最后
	var s6 string = "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" +
		"fmt.Println(flag01, unsafe.Sizeof(flag01))" +
		"abc" + "def" + "abc" + "def"
	fmt.Println(s6)
}
