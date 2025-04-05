package main

import (
	"fmt"
	"unsafe"
)

/*
整数类型
*/
func main() {
	{
		// 有符号整型，占用字节数：1, 2, 4, 8
		var a int8 = 127
		fmt.Println(a, a+1, unsafe.Sizeof(a))
		var b int16 = 32767
		fmt.Println(b, b+1, unsafe.Sizeof(b))
		var c int32 = 2147483647
		fmt.Println(c, c+1, unsafe.Sizeof(c))
		var d int64 = 9223372036854775807
		fmt.Println(d, d+1, unsafe.Sizeof(d))
		fmt.Println("------------------------------------------------------------------")
	}

	{
		// 无符号整型，占用字节数：1, 2, 4, 8
		var a uint8 = 255
		fmt.Println(a, a+1, unsafe.Sizeof(a))
		var b uint16 = 65535
		fmt.Println(b, b+1, unsafe.Sizeof(b))
		var c uint32 = 4294967295
		fmt.Println(c, c+1, unsafe.Sizeof(c))
		var d uint64 = 18446744073709551615
		fmt.Println(d, d+1, unsafe.Sizeof(d))
		fmt.Println("------------------------------------------------------------------")
	}

	{
		// 范围 [-128, 127]
		// cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)
		//var age int8 = 128
		//fmt.Println(age)
	}

}
