package main

import (
	"errors"
	"fmt"
)

/*
go中追求代码优雅，引入机制：defer+recover机制处理错误
*/
func main() {
	//deferAndRecover(2, 0)
	err := defineError(2, 0)
	if err != nil {
		fmt.Println("error =", err)
	}
	// panic终止程序
	if err != nil {
		panic(err)
	}

	fmt.Println("main函数成功结束")
}

// 测试panic

// defer+recover机制处理错误
func deferAndRecover(n1 int, n2 int) {
	defer dealError()
	res := n1 / n2
	fmt.Println(res)
}

// 自定义异常
func defineError(n1 int, n2 int) error {
	if n2 == 0 {
		return errors.New("除数不能是0")
	}
	res := n1 / n2
	fmt.Println(res)
	return nil
}

func dealError() {
	err := recover()
	if err != nil {
		fmt.Println("出现错误", err)
	}
}
