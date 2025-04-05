package main

import "fmt"

var num = initNum()

func initNum() int {
	fmt.Println("init global variable")
	return 11
}

func main() {
	fmt.Println("main..")
	fmt.Println("--------------------------------")

	// 匿名函数
	res := anonymous(2, 3, func(a int, b int) int {
		return a * b
	})
	fmt.Println(res)
	fmt.Println("--------------------------------")

	// 测试闭包
	f0 := myIncrease()
	fmt.Println(f0(1))
	fmt.Println(f0(2))
	fmt.Println(f0(3))
	fmt.Println(f0(4))
	fmt.Println("--------------------------------")

	// 在函数中，程序员经常需要创建资源，为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer关键字
	println(testDefer(1, 2))
}

func testDefer(n1 int, n2 int) int {
	defer fmt.Println("n1 =", n1)
	defer fmt.Println("n2 =", n2)
	sum := n1 + n2
	fmt.Println("sum =", sum)
	return sum
}

// init函数：初始化函数，可以用来进行一些初始化的操作
func init() {
	fmt.Println("init..")
}

func anonymous(n1 int, n2 int, myFunc func(int, int) int) int {
	return myFunc(n1, n2)
}

func myIncrease() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}
