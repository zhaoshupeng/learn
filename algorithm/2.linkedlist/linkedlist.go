package main

import "fmt"

/**
1、缓存淘汰策略来决定。常见的策略有三种：
	先进先出策略 FIFO（First In，First Out）、
	最少使用策略 LFU（Least Frequently Used）、
	最近最少使用策略 LRU（Least Recently Used）。

2、三种最常见的链表结构，它们分别是：单链表、双向链表和循环链表。

3、O(1): 针对链表的插入和删除操作，我们只需要考虑相邻结点的指针改变
	O(N):链表随机访问的性能没有数组好，需要 O(n) 的时间复杂度。
4、
 (1) 单链反转: 头结点插入法
 (2) 就地反转

5、链表常用技巧:
 (1) 快慢指针: 快指针的步长可以调整，一般为2;也可以和慢指针一样，比如删除倒数第n节点
 (2) 复制和判断时是否需要next,需慎重考虑

6、链表代码是否正确验证技巧:
	如果链表为空时，代码是否能正常工作？
	如果链表只包含一个结点时，代码是否能正常工作？
	如果链表只包含两个结点时，代码是否能正常工作？
	代码逻辑在处理头结点和尾结点的时候，是否能正常工作？
7、常见链表操作
	(1) 单链表反转
	(2) 链表中环的检测
    (3) 两个有序的链表合并
	(4) 删除链表倒数第 n 个结点
	(5) 求链表的中间结点

*/

func main() {
	singleList := NewSingleLinkedList()
	for i := 0; i < 9; i++ {

		singleList.InsertToTail(i)
	}
	singleList.Print()
	singleList.Reverse() //单链反转
	singleList.Print()

	Palindrome1 := NewSingleLinkedList()
	Palindrome1.InsertToTail("2")
	Palindrome1.InsertToTail("4")
	Palindrome1.InsertToTail("4")
	Palindrome1.InsertToTail("2")

	Palindrome2 := NewSingleLinkedList()
	Palindrome2.InsertToTail("2")
	Palindrome2.InsertToTail("4")
	Palindrome2.InsertToTail("5")
	Palindrome2.InsertToTail("4")
	Palindrome2.InsertToTail("2")
	fmt.Println(isPalindrome2(Palindrome1), isPalindrome2(Palindrome2))
}

type SingleLinkedListNode struct {
	value interface{}
	next  *SingleLinkedListNode
}
type SingleLinkedList struct {
	head   *SingleLinkedListNode
	length uint
}

func NewSingleLinkedListNode(v interface{}) *SingleLinkedListNode {
	return &SingleLinkedListNode{v, nil}
}

func (slln *SingleLinkedListNode) GetNext() *SingleLinkedListNode {
	return slln.next
}

func (slln *SingleLinkedListNode) GetValue() interface{} {
	return slln.value
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{NewSingleLinkedListNode(0), 0}
}

//在某个节点后面插入节点
func (sll *SingleLinkedList) InsertAfter(p *SingleLinkedListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewSingleLinkedListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	sll.length++
	return true
}

//在某个节点前面插入节点
func (sll *SingleLinkedList) InsertBefore(p *SingleLinkedListNode, v interface{}) bool {
	if p == nil || sll.head == p {
		return false
	}
	cur := sll.head.next // 目标节点
	pre := sll.head      // 目标节点的前继节点
	for cur != nil {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 目标节点未找到
	if cur == nil {
		return false
	}
	newNode := NewSingleLinkedListNode(v)
	pre.next = newNode
	newNode.next = cur
	sll.length++
	return true
}

//在链表头部插入节点
func (sll *SingleLinkedList) InsertToHead(v interface{}) bool {
	return sll.InsertAfter(sll.head, v)
}

//通过索引查找节点
func (sll *SingleLinkedList) InsertToTail(v interface{}) bool {
	cur := sll.head
	for cur.next != nil {
		cur = cur.next
	}
	return sll.InsertAfter(cur, v)
}

// 删除节点
func (sll *SingleLinkedList) DeleteNode(p *SingleLinkedListNode) bool {
	if p == nil {
		return false
	}
	cur := sll.head.next
	pre := sll.head

	for cur != nil { // 未到尾结点
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil { // 未找到目标节点
		return false
	}

	pre.next = cur.next
	p = nil //todo ???
	sll.length--
	return true
}

//打印链表
//时间复杂度：O(N)
func (sll *SingleLinkedList) Print() {
	cur := sll.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// (1) 单链表反转
func (sll *SingleLinkedList) Reverse() {
	if sll.head == nil || sll.head.next == nil || sll.head.next.next == nil {
		return
	}
	var pre *SingleLinkedListNode = nil // 待处理节点的后继节点。初始为空
	cur := sll.head.next
	// 从链表第一个节点开始反转
	for cur != nil {
		tmp := cur.next // 待反转的下一个节点
		cur.next = pre
		pre = cur
		cur = tmp
	}
	sll.head.next = pre
}

/*
判断单链表是否有环
*/
// (2) 链表中环的检测
func (sll *SingleLinkedList) HasCycle() bool {
	if sll.head != nil {
		slow := sll.head
		fast := sll.head

		// 只要有环，快指针总能追上慢指针
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
(3) 两个有序单链表合并
*/
func MergeSortedList(l1, l2 *SingleLinkedList) *SingleLinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}

	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	l := NewSingleLinkedList()
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next

	for curl1 != nil && curl2 != nil {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}

	}

	if curl1 == nil {
		cur.next = curl2
	} else {
		cur.next = curl1
	}
	l.length = l1.length + l2.length

	return l
}

/*
(4)删除倒数第N个节点
*/
func (sll *SingleLinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == sll.head || nil == sll.head.next {
		return
	}

	// 先让前进指针提前一定长度
	fast := sll.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}
	slow := sll.head // 最终slow的下一个节点就是待删除的节点
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

/*
(5) 获取中间节点
*/
func (sll *SingleLinkedList) FindMiddleNode() *SingleLinkedListNode {
	if nil == sll.head || nil == sll.head.next {
		return nil
	}
	// 就存在一个真实有效节点
	if nil == sll.head.next.next {
		return sll.head.next
	}

	slow, fast := sll.head, sll.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/
func isPalindrome2(l *SingleLinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	var isPalindrome = true
	step := lLen / 2
	var pre *SingleLinkedListNode = nil
	cur := l.head.next
	next := l.head.next.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	mid := cur
	//fmt.Println(pre.value, mid.value)  // 如果是偶数个节点此时中间节点是靠后的节点

	var left, right *SingleLinkedListNode = pre, nil
	if lLen%2 != 0 { //兼容节点个数是奇数的情况
		right = mid.GetNext()
	} else {
		right = mid
	}

	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre

	return isPalindrome
}
