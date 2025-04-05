package main

import (
	"fmt"
)

/*
1. Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存，超出1字节使用int
2. Golang中字符使用UTF-8编码
*/
func main() {
	//定义字符类型的数据：
	var c1 byte = 'a'
	fmt.Println(c1) //97
	var c2 byte = '6'
	fmt.Println(c2)   //54
	var c3 byte = '(' // ASCII是40
	fmt.Println(c3 + 20)
	//字符类型，本质上就是一个整数，也可以直接参与运算，输出字符的时候，会将对应的码值做一个输出
	//字母，数字，标点等字符，底层是按照ASCII进行存储。
	var c4 int = '中'
	fmt.Println(c4)
	//汉字字符，底层对应的是Unicode码值
	//对应的码值为20013，byte类型溢出，可以用int
	//总结：Golang的字符对应的使用的是UTF-8编码（Unicode是对应的字符集，UTF-8是Unicode的其中的一种编码方案）
	var c5 byte = 'A'
	//想显示对应的字符，必须采用格式化输出
	fmt.Printf("c5对应的具体的字符为：%c\n", c5)

	fmt.Println("------------------------------------------------------------------")
	escape()
}

// 练习转义字符：
func escape() {
	//\n  换行
	fmt.Println("aaa\nbbb")
	//\b 退格
	fmt.Println("aa啊\bbbb")
	fmt.Println("------------------------")
	//\r 回车：光标回到本行的开头
	fmt.Println("aaaaa啊啊\rbbb")
	fmt.Println("------------------------")
	//\t 制表符
	fmt.Println("aaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("aaaaaaaa\tbbbbb")
	//\"
	fmt.Println("\"Golang\"")
}
