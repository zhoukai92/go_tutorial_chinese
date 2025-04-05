package main

import "fmt"

func main() {
	// for循环的语法格式：
	// for 初始表达式; 布尔表达式（条件判断）; 迭代因子 {
	// 	循环体;--》反复重复执行的内容
	// }
	// 注意：for的初始表达式 不能用var定义变量的形式，要用:=
	// 注意：for循环实际就是让程序员写代码的效率高了，但是底层该怎么执行还是怎么执行的，底层效率没有提高，只是程序员写代码简洁了而已
	//实现一个功能：求和： 1 + 2 + ... + 100 = ?
	var sum int
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 只有条件
	//onlyCondition()

	// 死循环
	//endlessLoop()
	// 演示break
	//demoBreak()
	//demoBreak1()
	demoContinue()
}

func demoContinue() {
	//功能：输出1-100中被23整除的数：
	n := 23
	for i := 0; i < 100; i++ {
		if i%n == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("-------------")
	// 方式二
	for i := 0; i < 100; i++ {
		if i%n != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func demoBreak1() {
	//双重循环：
outer:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v \n", i, j)
			if i == 2 && j == 2 {
				break outer //结束指定标签对应的循环
			}
		}
	}
	fmt.Println("-----ok")
}

func demoBreak() {
	//功能：求1-100的和，当和第一次超过300的时候，停止程序
	var sum = 0
	i := 0
	for ; i < 100; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			//停止正在执行的这个循环：
			break
		}
	}
	fmt.Println("-----ok", i)
}

func endlessLoop() {
	// 在控制台中结束死循环：ctrl+c
	for {
		fmt.Println("hello world")
	}
}

func onlyCondition() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}
}
