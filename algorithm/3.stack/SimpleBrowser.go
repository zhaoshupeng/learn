package main

import "fmt"

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	//Top() interface{}
	Flush()
}

type Browser struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

func (b *Browser) IsCanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) IsCanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forwardStack.Flush()
}

func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}