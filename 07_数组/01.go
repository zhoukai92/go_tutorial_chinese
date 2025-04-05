package main

import (
	"fmt"
)

/*
二维数组
*/
func main() {
	var arr = [2][3]int16{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr)

	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[1][2])
	fmt.Println("---------------------------------------------------")

	// 遍历
	for i := 0; i < len(arr); i++ {
		nums := arr[i]
		for j := 0; j < len(nums); j++ {
			fmt.Printf("a[%d][%d] = %v\t", i, j, nums[j])
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------------------")
	for i, nums := range arr {
		for j, num := range nums {
			fmt.Printf("a[%d][%d] = %v\t", i, j, num)
		}
		fmt.Println()
	}
}
