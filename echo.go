package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int64 = 10
	fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p\n", num, unsafe.Sizeof(num), &num)

	var sli []interface{}
	fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p; 变量值:%+v\n", sli, unsafe.Sizeof(sli), &sli, sli)
	fmt.Println(sli)
	fmt.Println(gcd(8, 3))
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
