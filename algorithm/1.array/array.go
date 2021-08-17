package main

import (
	"errors"
	"fmt"
)

/**
概括：
数组：连续的内存空间和相同类型的数据。与链表相比，支持随机访问。
	a. 低效的“插入”和“删除”：(数组有序情况下)插入和删除的最坏复杂度是O(n),每个位置插入的概率一样，所以平均时间复杂度(1+2+...n)/n=O(n)
	b.为了避免搬移数据，可以每次的删除操作并不是真正地搬移数据，只是记录数据已经被删除。当数组没有更多空间存储数据时，我们再触发执行一次真正的删除操作
	c.警惕数组的访问越界问题.
线性表：数组、链表、队列、栈。(线性表上数据最多只有前和后两个方向)

*/

/** 问题：


* 1) 为什么数组要从 0 开始编号，而不是从 1 开始呢？
	从 1 开始编号，每次随机访问数组元素都多了一次减法运算，对于 CPU 来说，就是多了一次减法指令。
* 2) 容器能否完全替代数组？
		ArrayList 最大的优势就是可以将很多数组操作的细节封装起来.并且支持动态扩容。
* 3) 二维数组寻址公式：
	对于 m * n 的数组，a [ i ][ j ] (i < m,j < n)的地址为：
	address = base_address + ( i * n + j) * type_size

*/

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2) 数组中的数据是int类型的；
 *
 *
 */

type Array struct {
	// slice三个属性：指针(指针指向数组的第一个可以从slice中访问的元素)、长度(元素的个数)和容量(通常是从slice的起始元素到底层数组的最后一个元素的个数据)。
	data   []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (arr *Array) Len() uint {
	return arr.length
}

//判断索引是否越界
func (arr *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(arr.data)) {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (arr *Array) Find(index uint) (int, error) {
	if arr.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return arr.data[index], nil
}

func (arr *Array) Insert(index uint, v int) error {
	// 单独一个变量标记数组长度
	if arr.Len() == uint(cap(arr.data)) {
		return errors.New("full array")
	}
	if index != arr.length && arr.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	// 移动数据
	for i := arr.length; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = v
	arr.length++
	return nil
}

func (arr *Array) Delete(index uint) (int, error) {
	if arr.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := arr.data[index]
	for i := index; i < arr.Len()-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.length--
	return v, nil
}

func main() {
	arr := NewArray(2)
	err := arr.Insert(uint(0), 2)
	fmt.Println(err, arr)
	err = arr.Insert(uint(0), 1)
	fmt.Println(err, arr)
	//err = arr.Insert(uint(1), 1)
	//fmt.Println(err, arr)
	//err = arr.Insert(uint(2), 2)
	//fmt.Println(err, arr)
}
