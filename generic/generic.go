package main

import (
	"fmt"
)

// 泛型
// https://www.liwenzhou.com/posts/Go/generics/
//https://www.cnblogs.com/apocelipes/p/17576990.html
//https://cloud.tencent.com/developer/article/2029500

func main() {
	sum := add[int](3, 5)
	fmt.Println("add: ", sum)

	fmin := min[float64] // 类型实例化，编译器生成T=float64的min函数
	m2 := fmin(1.2, 2.3) // 1.2
	m3 := min[int](9, 6)
	fmt.Println("min: ", m2, m3)

	mp := map[int]string{1: "good"}
	fmt.Println("map: ", mp)

	PrintOrdered[Int](Int(9))
	PrintOrdered[IntItem](IntItem{9})

}

// 1.
// 类型形参和类型实参
func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

//func add[T int | string](a, b T) T {
//	return a + b
//}

// Go语言中的类型约束是接口类型。类型约束接口可以直接在类型参数列表中使用。
// 类型约束字面量，通常外层interface{}可省略
func add[T interface{ int | float64 }](a, b T) T {
	return a + b
}

// Go1.18开始接口类型的定义也发生了改变，由过去的接口类型定义方法集（method set）变成了接口类型定义类型集（type set）。**也就是说，接口类型现在可以用作值的类型，也可以用作类型约束。
//从 Go 1.18 开始，一个接口不仅可以嵌入其他接口，还可以嵌入任何类型、类型的联合或共享相同底层类型的无限类型集合.
// 接口作为类型集是一种强大的新机制，是使类型约束能够生效的关键。目前，使用新语法表的接口只能用作类型约束。

// 作为类型约束使用的接口类型可以事先定义并支持复用。
// 事先定义好的类型约束类型
type Value interface {
	int | float64
}

func min2[T Value](a, b T) T {
	return a + b
}

// 在使用类型约束时，如果省略了外层的interface{}会引起歧义，那么就不能省略。例如
// type IntPtrSlice1 [T *int][]T             // T*int ?
type IntPtrSlice2[T *int,] []T             // 只有一个类型约束时可以添加`,`
type IntPtrSlice3[T interface{ *int }] []T // 使用interface{}包裹

// 2.
// 除了函数中支持使用类型参数列表外，类型也可以使用类型参数列表。
type Slice[T int | string] []T

type Map[K int | string, V float32 | float64] map[K]V

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

type Int int

type IntItem struct {
	V int
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | IntItem |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string // ~X 表示支持 X 的衍生类型
	String() string
}

func (i Int) String() string {
	return "string value"
}

func (i IntItem) String() string {
	return "string value"
}

func PrintOrdered[T Ordered](v T) {
	fmt.Printf("------PrintOrdered: %v \n", v)
	fmt.Println("v  string value: ", v.String())
}

type Iterator[T any] interface {
	ForEachRemaining(action func(T) error) error
	// other methods
}
