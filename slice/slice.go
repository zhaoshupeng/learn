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

	//s1 := []int{0, 1, 2, 3, 8: 100}
	//fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p; 变量值:%+v\n", s1, unsafe.Sizeof(s1), &s1, s1)
	//fmt.Println("切片长度:", len(s1), " 切片容量:", cap(s1))

	// nil 切片
	var s1 []int
	// var s1 = *new([]int)	//// new 函数返回是指针类型，所以需要使用 * 号来解引用
	fmt.Println("nil切片, 长度：", len(s1), " 容量：", cap(s1), " 与nil比较: ", s1 == nil)

	// 空 切片
	var s2 = []int{}
	// var s2 = make([]int,0)
	fmt.Println("空切片, 长度：", len(s2), " 容量：", cap(s2), " 与nil比较: ", s2 == nil)

	fmt.Println("------------------------------------------")
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[5:7:9] // data[low, high, max]
	//  high 和 max 则是开区间，表示最后一个元素只能是索引 high-1 处的元素，而最大容量则只能是索引 max-1 处的元素。
	// 当 high == low 时，新 slice 为空。
	// 还有一点，high 和 max 必须在老数组或者老 slice 的容量（cap）范围内。
	//fmt.Println(slice, data, len(slice), cap(slice), &slice)
	fmt.Printf("11  slice变量值:%+v; data变量值:%+v; slice长度:%d; slice容量:%d; slice地址：%p;slice第一个值的地址:%p\n", slice, data, len(slice), cap(slice), &slice, unsafe.Pointer(&slice[0]))
	slice = append(slice, 77)
	fmt.Printf("22  slice变量值:%+v; data变量值:%+v; slice 长度:%d; slice 容量:%d; slice地址：%p;slice第一个值的地址:%p\n", slice, data, len(slice), cap(slice), &slice, unsafe.Pointer(&slice[0]))

	slice = append(slice, 88)
	fmt.Printf("33  slice变量值:%+v; data变量值:%+v; slice 长度:%d; slice 容量:%d; slice地址：%p;slice第一个值的地址:%p\n", slice, data, len(slice), cap(slice), &slice, unsafe.Pointer(&slice[0]))

	slice = append(slice, 99)
	fmt.Printf("44  slice变量值:%+v; data变量值:%+v; slice 长度:%d; slice 容量:%d; slice地址：%p;slice第一个值的地址:%p\n", slice, data, len(slice), cap(slice), &slice, unsafe.Pointer(&slice[0]))

	slice = append(slice, 10)
	fmt.Printf("55  slice变量值:%+v; data变量值:%+v; slice 长度:%d; slice 容量:%d; slice地址：%p;slice第一个值的地址:%p\n", slice, data, len(slice), cap(slice), &slice, unsafe.Pointer(&slice[0]))

	fmt.Println("------------------------------------------")
	var s11 []int
	var s22 = []int{}
	var s3 = make([]int, 0)
	var s4 = *new([]int)

	var a1 = *(*[3]int)(unsafe.Pointer(&s11))
	var a2 = *(*[3]int)(unsafe.Pointer(&s22))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test := [3]int{1, 2, 3}
	testSli := test[0:2]
	testSliT := test[0:2]
	testSliP := &testSliT
	fmt.Printf("test:%p,test0:%p\n", &test, &test[0])
	fmt.Printf("test:%p,test0:%p\n", &testSli, &testSli[0])
	fmt.Printf("testSliP:%p,testSliP0:%p\n", testSliP, &(*testSliP)[0])

}
