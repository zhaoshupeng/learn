package main

//func main() {
//	arrayStack := NewArrayStack()
//	arrayStack.Push(5)
//	arrayStack.Push(4)
//	arrayStack.Push(3)
//	fmt.Println(*arrayStack)
//	arrayStack.Pop()
//	arrayStack.Pop()
//	fmt.Println(*arrayStack)
//	arrayStack.Push(6)
//
//	fmt.Println(*arrayStack)
//
//}

/*
基于数组实现的栈: 顺序栈
*/
type ArrayStack struct {
	// 数据
	data []interface{}
	// 栈顶位置
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 2),
		top:  -1,
	}
}

// 入栈
func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top += 1
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

// 出栈
func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	v := as.data[as.top]
	as.top -= 1
	return v
}

func (as *ArrayStack) IsEmpty() bool {
	if as.top < 0 {
		return true
	}
	return false
}

func (as *ArrayStack) Flush() {
	as.top = -1
}
