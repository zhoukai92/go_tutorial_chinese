package main

import "fmt"

func main() {
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：
	//定义一个数组：
	var scores [5]int
	//将成绩存入数组：
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21
	//求和：
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	//平均数：
	avg := sum / len(scores)
	//输出
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v\n", sum, avg)
	// range 遍历
	for i, score := range scores {
		fmt.Println(i, score)
	}
	fmt.Println("---------------------------------------------------")

	// 数组元素地址
	{
		var arr [3]int16
		//获取数组的长度：
		fmt.Println(len(arr), arr)
		//证明arr中存储的是地址值：
		fmt.Printf("arr的地址为：%p\n", &arr)
		//第一个空间的地址：
		fmt.Printf("arr的地址为：%p\n", &arr[0])
		//第二个空间的地址：
		fmt.Printf("arr的地址为：%p\n", &arr[1])
		//第三个空间的地址：
		fmt.Printf("arr的地址为：%p\n", &arr[2])
		fmt.Println("---------------------------------------------------")
	}

	// 初始化的方式
	// 长度是类型的一部分
	{
		//第一种：
		var arr1 [3]int = [3]int{3, 6, 9}
		fmt.Printf("%v, type = %T\n", arr1, arr1)
		//第二种：
		var arr2 = [3]int{1, 4, 7}
		fmt.Printf("%v, type = %T\n", arr2, arr2)
		//第三种：
		var arr3 = [...]int{4, 5, 6, 7}
		fmt.Printf("%v, type = %T\n", arr3, arr3)
		//第四种：
		var arr4 = [...]int{2: 66, 0: 33, 8: 88}
		fmt.Printf("%v, type = %T\n", arr4, arr4)
		fmt.Println("---------------------------------------------------")
	}
	//Go中数组属值类型，在默认情况下是值传递，因此会进行值拷贝。
	nums := [3]int{0, 1, 2}
	test00(nums)
	fmt.Println(nums)
	// 使用指针（引用传递）
	test01(&nums)
	fmt.Println(nums)

}

func test00(nums [3]int) {
	nums[0] = 10
}

func test01(nums *[3]int) {
	nums[0] = 10
}
