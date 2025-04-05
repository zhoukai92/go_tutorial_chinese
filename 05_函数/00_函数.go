package main

import "fmt"

//	func   函数名（形参列表)（返回值类型列表）{
//		执行语句..
//		return + 返回值列表
//	}
func main() {
	fmt.Println(mySum(10, 20))
	sum, sub := sumAndSub(10, 20)
	fmt.Println(sum, sub)
	fmt.Println(sumAndSub2(10, 20))
	fmt.Println("--------------------------------")

	someNum(1, 2, 3)
	fmt.Println("--------------------------------")

	a := 2
	increase(&a)
	fmt.Println(a)
	fmt.Println("--------------------------------")

	// 函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	sumFun := mySum
	fmt.Println(sumFun(2, 3))
	fmt.Printf("sumFun 的类型是%T\n", sumFun)
	fmt.Println("--------------------------------")
	// 函数作为入参
	fmt.Println(calculate(1, 2, mySum))
	fmt.Println(calculate(1, 2, mySub))
}

func sumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func sumAndSub2(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 自定义函数：功能：两个数相加：
func mySum(n1 int, n2 int) int { //如果返回值类型就一个的话，那么()是可以省略不写的
	var sum int = 0
	sum += n1
	sum += n2
	return sum
}
func mySub(n1 int, n2 int) int { //如果返回值类型就一个的话，那么()是可以省略不写的
	return n1 - n2
}

//Golang中函数不支持重载

// Golang中支持可变参数
// 定义一个函数，函数的参数为：可变参数 ...  参数的数量可变
// args...int 可以传入任意多个数量的int类型的数据  传入0个，1个，，，，n个
func someNum(args ...int) {
	//函数内部处理可变参数的时候，将可变参数当做切片来处理
	//遍历可变参数：
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

// 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
/*
以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，
可以传入变量的地址&，函数内以指针的方式操作变量。从效果来看类似引用传递。
*/

func increase(num *int) {
	*num = *num + 1
}

func calculate(n1 int, n2 int, myFunc func(int, int) int) int {
	return myFunc(n1, n2)
}
