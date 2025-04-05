package main

import (
	"fmt"
	"strconv"
)

/*
基本数据类型间的转换
1. Go在不同类型的变量之间赋值时需要显式转换，并且只有显式转换(强制转换)。
2. 语法： 表达式T(v)将值v转换为类型T
*/
func main() {
	// 整型间的转化
	int2int()
	fmt.Println("------------------------------------------------------------------------------")
	// 转化为字符串
	other2string()
	fmt.Println("------------------------------------------------------------------------------")
	// string转化为其他类型
	string2other()
}

// string转化为其他类型
func string2other() {
	//string-->bool
	var s1 = "true"
	var b bool
	//ParseBool这个函数的返回值有两个：(value bool, err error)
	//value就是我们得到的布尔类型的数据，err出现的错误
	//我们只关注得到的布尔类型的数据，err可以用_直接忽略
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b的类型是：%T,b=%v \n", b, b)

	//string---》int64
	var s2 = "19"
	var num1 int64
	num1, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num1的类型是：%T,num1=%v \n", num1, num1)

	//string-->float32/float64
	var s3 = "3.14"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("f1的类型是：%T,f1=%v \n", f1, f1)
	fmt.Println("-------------")

	//注意：string向基本数据类型转换的时候，一定要确保string类型能够转成有效的数据类型，
	//否则最后得到的结果就是按照对应类型的默认值输出
	var s4 = "golang"
	var b1, err = strconv.ParseBool(s4)
	fmt.Printf("b1的类型是：%T,b1=%v \n", b1, b1)
	fmt.Println(err)

	var s5 = "golang"
	var num2, err2 = strconv.ParseInt(s5, 10, 64)
	fmt.Printf("num2的类型是：%T,num2=%v \n", num2, num2)
	fmt.Println(err2)
}

// 转化为字符串
func other2string() {
	var n1 = 19
	var f2 float32 = 4.78
	var b3 = false
	var b4 byte = 'a'
	var s1 = fmt.Sprintf("%d", n1)
	fmt.Printf("s1对应的类型是：%T ，s1 = %q \n", s1, s1)
	var s2 = fmt.Sprintf("%f", f2)
	fmt.Printf("s2对应的类型是：%T ，s2 = %q \n", s2, s2)
	var s3 = fmt.Sprintf("%t", b3)
	fmt.Printf("s3对应的类型是：%T ，s3 = %q \n", s3, s3)
	var s4 = fmt.Sprintf("%c", b4)
	fmt.Printf("s4对应的类型是：%T ，s4 = %q \n", s4, s4)
}

// 整型间的转化
func int2int() {
	//进行类型转换：
	var n1 = 100
	//var n2 float32 = n1  // 在这里自动转换不好使，比如显式转换
	fmt.Println(n1)
	//fmt.Println(n2)
	var n2 = float32(n1)
	fmt.Println(n2)
	//注意：n1的类型其实还是int类型，只是将n1的值100转为了float32而已，n1还是int的类型
	fmt.Printf("%T", n1) //int
	fmt.Println("\n--------")
	//将int64转为int8的时候，编译不会出错的，但是会数据的溢出
	var n3 int64 = 888888
	var n4 = int8(n3)
	fmt.Println(n4) //56
	var n5 int32 = 12
	var n6 = int64(n5) + 30 //一定要匹配=左右的数据类型
	fmt.Println(n5)
	fmt.Println(n6)
	var n7 int64 = 12
	var n8 = int8(n7) + 127 //编译通过，但是结果可能会溢出
	fmt.Println(n8)
	//var n9 int8 = int8(n7) + 128 //编译不会通过
	//fmt.Println(n9)
}
