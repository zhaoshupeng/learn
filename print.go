package main

import (
	"fmt"
	"runtime"
	"time"
)

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		//recover()
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}
func main1() {

	score := float64(time.Now().Unix())
	fmt.Println("---------------: ", score)

	//var root *BinaryTreeNode
	////root = &BinaryTreeNode{
	////	Data: 11,
	////}
	//var t *BinaryTreeNode
	//t = &BinaryTreeNode{
	//	Data: 33,
	//}
	//fmt.Printf("t地址：%p，t指针的值：%p, t的值的默认格式：%v t的值的默认格式：%+v \n", &t, t, t, t)
	//fmt.Println(unsafe.Pointer(t))
	//
	//create(root)
	//fmt.Println("-----------------------")
	//fmt.Printf("root地址：%p，值：%+v\n", root, root)
	//
	////f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0777) //读写模式打开，写入追加
	////defer f.Close()
	////add_data := "this is add"
	////num, _ := f.Write([]byte(add_data))
	////fmt.Println(num)
	////if test() == nil {
	////
	////}
	////fmt.Println("ssss", test() == nil)
	////fmt.Println("ssss", test(), test1())
}

func test() []interface{} {
	return nil
}
func test1() []int {
	return nil
}

type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func create(node *BinaryTreeNode) {
	//node = &BinaryTreeNode{
	//	Data: 2,
	//}
	fmt.Println("参数：", node)
	tem := &BinaryTreeNode{
		Data: 2,
	}
	fmt.Println("变量tem：", tem)
	fmt.Printf("变量值：%v", tem)
	fmt.Printf("值： %v, 类型： %T", tem, tem)

	node = tem
	//node.Data = 55
	fmt.Println("node函数内: ", node)
	fmt.Printf("node: %v; p : %p\n", node, node)
}

const (
	REDIS = "redis"
	MYSQL = "mysql"
)
