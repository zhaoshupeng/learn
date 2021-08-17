package main

import "fmt"

type node struct {
	val  interface{}
	next *node
}

type LinkedListStack struct {
	//栈顶节点
	top *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (lls *LinkedListStack) IsEmpty() bool {
	return lls.top == nil
}

func (lls *LinkedListStack) Push(v interface{}) {
	lls.top = &node{next: lls.top, val: v}
}

func (lls *LinkedListStack) Pop() interface{} {
	if lls.IsEmpty() {
		return nil
	}

	v := lls.top.val
	lls.top = lls.top.next
	return v
}

func (lls *LinkedListStack) Flush() {
	lls.top = nil
}

func (lls *LinkedListStack) Print() {
	if lls.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := lls.top
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
