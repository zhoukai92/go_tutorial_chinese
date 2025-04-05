package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 字符串函数
	//testString()

	// 时间函数
	testTime()
}

func testTime() {
	//时间和日期的函数，需要到入time包，所以你获取当前时间，就要调用函数Now函数：
	now := time.Now()
	//Now()返回值是一个结构体，类型是：time.Time
	fmt.Printf("%v ~~~ 对应的类型为：%T\n", now, now)
	//2021-02-08 17:47:21.7600788 +0800 CST m=+0.005983901 ~~~ 对应的类型为：time.Time
	//调用结构体中的方法：
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", now.Month())      //月：February
	fmt.Printf("月：%v \n", int(now.Month())) //月：2
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())
	fmt.Printf("UnixMilli：%v \n", now.UnixMilli())
	fmt.Printf("UnixMicro：%v \n", now.UnixMicro())
	fmt.Printf("UnixNano：%v \n", now.UnixNano())
	fmt.Printf("Nanosecond：%v \n", now.Nanosecond())

	fmt.Println("--------------------------------")
	//Printf将字符串直接输出：
	fmt.Printf("当前时间： %d-%02d-%02d %02d:%02d:%02d  \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	//Sprintf可以得到这个字符串，以便后续使用：
	datestr := fmt.Sprintf("当前时间： %d-%02d-%02d %02d:%02d:%02d  \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)
	fmt.Println("--------------------------------")

	//这个参数字符串的各个数字必须是固定的，必须这样写
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)
	//选择任意的组合都是可以的，根据需求自己选择就可以（自己任意组合）。
	datestr3 := now.Format("2006 15:04")
	fmt.Println(datestr3)
}

func testString() {
	golang := "golang 你好"
	fmt.Println(len(golang)) // 字节个数
	fmt.Println("--------------------------------")

	// for-range键值循环：
	for i, c := range golang {
		fmt.Println(i, c, string(c))
	}
	fmt.Println("--------------------------------")

	r := []rune(golang)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i], string(r[i]))
	}
	fmt.Println("--------------------------------")

	n, err := strconv.Atoi("123")
	if err != nil {
		return
	}
	fmt.Println(n)
	fmt.Println(strconv.Itoa(123))
	fmt.Println("--------------------------------")

	fmt.Println("是否包含", strings.Contains(golang, "go"))
	fmt.Println("计数", strings.Count(golang, "go"))
	fmt.Println("不区分大小写比较", strings.EqualFold("GO", "go"))
	fmt.Println("搜索子串", strings.Index(golang, "o"))
	fmt.Println("替换", strings.Replace("go1go2go3", "go", "xx", 2))
	fmt.Println(strings.Split("c-java-go-python-js", "-"))
	fmt.Println("--------------------------------")
	fmt.Println(strings.ToUpper(golang), strings.ToLower(golang))
	fmt.Println(strings.TrimSpace("hello world"))
	fmt.Println(strings.Trim("----golang-", "-"))
	fmt.Println(strings.HasPrefix("http://studygolang.com/pkgdoc", "http"))
}
