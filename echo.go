package main

import (
	"fmt"
	"hash/crc32"
	"unsafe"
)

func main() {

	// 创建一个新的链表，并向链表里面添加几个数字
	//l := list.New()
	//fmt.Printf("ddddd%#v\n", l)
	//fmt.Printf("%p\n", l)
	//
	//ll := new(list.List)
	//fmt.Printf("fffff%#v\n", ll)
	//fmt.Printf("%p\n", ll)
	var u User
	fmt.Println("u1 size is ", unsafe.Sizeof(u))

	re := new(Re)
	fmt.Printf("fffff%#v\n", re)
	fmt.Printf("%p\n", re)
	fmt.Printf("%p\n", &re.CourseId)
	fmt.Printf("%p\n", &re.Base)

	//e4 := l.PushBack(4)
	//e1 := l.PushFront(1)
	//l.InsertBefore(3, e4)
	//l.InsertAfter(2, e1)

	//fmt.Println("wwww", e4)
	//md5Str("good")
	//hashValue("good")
	//buckets := make([][]int, 3)
	//fmt.Println(len(buckets), buckets)
	//
	//arr := []int{5, 6, 8, 9}
	//for val := range arr {
	//	fmt.Println("val----", val)
	//}
	//
	//var num int64 = 10
	//fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p\n", num, unsafe.Sizeof(num), &num)
	//
	//var sli []interface{}
	//fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p; 变量值:%+v\n", sli, unsafe.Sizeof(sli), &sli, sli)
	//fmt.Println(sli)
	//fmt.Println(gcd(8, 3))
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func hashValue(origin string) {
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(origin), crcTable)
	fmt.Println("hashValue-----------", hashVal)
}

type Re struct {
	CourseId int `json:"course_id"`
	Base
}

type Base struct {
	Grade   int `json:"grade"`
	Subject int `json:"subject"`
}

// 64位平台，对齐参数是8
type User struct {
	A int32 // 4
	//	B []int32 // 24
	C string // 16
	D bool   // 1
}
