package main

import "fmt"

func main() {
	//实现功能：根据给出的学生分数，判断学生的等级：
	// >=90  -----A
	// >=80  -----B
	// >=70  -----C
	// >=60  -----D
	// <60   -----E
	//定义一个学生的成绩：
	//var score = getScore()
	// 使用 if 实现
	//demoIf(score)
	// 使用 Switch 实现
	//demoSwitch(score)

	// Switch特殊说明
	//switch00()
	switchFallthrough(2)
}

func getScore() int {
	fmt.Println("输入一个整数")
	var score int
	n, err := fmt.Scanln(&score)
	if err != nil {
		fmt.Println("输入的不是整数", err, n)
	}
	return score
}

func demoSwitch(score int) {
	//switch后面是一个表达式，这个表达式的结果依次跟case进行比较，满足结果的话就执行冒号后面的代码。
	//default是用来“兜底”的一个分支，其它case分支都不走的情况下就会走default分支
	// default语句不是必须的，位置也是随意的。
	// case后面不需要带break
	switch score / 10 {
	case 10:
		fmt.Println("您的等级为A级")
	case 9:
		fmt.Println("您的等级为A级")
	case 8:
		fmt.Println("您的等级为B级")
	case 7:
		fmt.Println("您的等级为C级")
	case 6:
		fmt.Println("您的等级为D级")
	case 5:
		fmt.Println("您的等级为E级")
	case 4:
		fmt.Println("您的等级为E级")
	case 3:
		fmt.Println("您的等级为E级")
	case 2:
		fmt.Println("您的等级为E级")
	case 1:
		fmt.Println("您的等级为E级")
	case 0:
		fmt.Println("您的等级为E级")
	default:
		fmt.Println("您的成绩有误")
	}
}

func demoIf(score int) {
	if score >= 90 {
		fmt.Println("您的成绩为A级别")
	} else if score >= 80 { //else隐藏：score < 90
		fmt.Println("您的成绩为B级别")
	} else if score >= 70 { //score < 80
		fmt.Println("您的成绩为C级别")
	} else if score >= 60 { //score < 70
		fmt.Println("您的成绩为D级别")
	} else { //score < 60
		fmt.Println("您的成绩为E级别")
	}
}

// switch后也可以不带表达式，当做if分支来使用
func switch00() {
	a := 1
	switch {
	case a < 1:
		fmt.Println("a < 1")
	case a == 1:
		fmt.Println("a == 1")
	case a > 1:
		fmt.Println("a > 1")
	}
}

// switch穿透，利用fallthrough关键字，如果在case语句块后增加fallthrough ,则会继续执行下一个case,也叫switch穿透。
func switchFallthrough(num int) {
	switch num {
	case 1:
		fmt.Println("num is 1")
	case 2:
		fmt.Println("num is 2")
		fallthrough
	case 3:
		fmt.Println("num is 3")
	}
}
