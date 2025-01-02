package main

import "fmt"

// Comparable 定义一个接口用于比较两个元素
//type Comparable interface {
//	Less(other Comparable) bool
//}

//type Pair struct {
//	ID    string
//	Value int
//}

//type Pair struct {
//	ID    string
//	Value sort.Interface
//}

// 以小顶堆 为例
type SmallHeap struct {
	arr     []*Pair //存储堆的数组，从下标 1 开始存储数据
	num     int     // 堆可以存储的最大数据个数
	count   int     // 堆中已经存储的数据个数
	sort    bool    // 是否使用了堆排序
	sortNum int     //排序后的数组数量
}

// init heap
// 初始化堆/建堆，capacity
func NewSmallHeap(capacity int) *SmallHeap {
	heap := SmallHeap{}
	heap.arr = make([]*Pair, capacity+1) // 因为从下标1 开始存储数据
	heap.num = capacity
	heap.count = 0
	return &heap
}

/**
一个包含n个节点的完全二叉树，其树高不超过log2N，堆化的时间复杂度跟树高成正比，也就是O(logN),插入和删除堆顶元素的主要逻辑就是堆化
*/

// top-max heap -> heapify from down to up
// 新增一个数据
func (heap *SmallHeap) Insert(data *Pair) {
	//defensive
	if heap.count == heap.num { // 堆满了
		if heap.arr[1].Value < data.Value {
			heap.RemoveMin()
			heap.Insert(data)
		}
		return
	}
	heap.count++
	heap.arr[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.arr[parent].Value > heap.arr[i].Value { // 自下往上堆化
		//swapPair(heap.arr, i, parent) //交换数组中的两个值
		heap.swap(i, parent)
		i = parent
		parent = i / 2
	}
}

// swap two elements
//func swapPair(a []*Pair, i int, j int) {
//	tmp := a[i]
//	a[i] = a[j]
//	a[j] = tmp
//}

func (heap *SmallHeap) RemoveMin() {
	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	//删除堆顶元素，并将堆的最后的值放置到堆顶，然后从堆顶进行堆化，这样结束后还是能满足完全二叉树的特性(或堆的定义)
	heap.swap(1, heap.count)
	heap.count--

	//heapify from up to down
	// 从上向下堆化
	//heapifyUpToDownPairSmall(heap.arr, heap.count)
	heap.HeapifyUpToDown()
}

func (heap *SmallHeap) HeapifyUpToDown() {
	for i := 1; i <= heap.count/2; {
		maxIndex := i

		if heap.arr[i].Value > heap.arr[i*2].Value {
			maxIndex = i * 2
		}
		if i*2+1 <= heap.count && heap.arr[maxIndex].Value > heap.arr[i*2+1].Value {
			maxIndex = i*2 + 1
		}
		// 已经大于左右子树的任意值
		if maxIndex == i {
			break
		}
		//swapPair(arr, i, maxIndex)
		heap.swap(i, maxIndex)
		i = maxIndex
	}
	return
}
func (heap *SmallHeap) swap(i, j int) {
	tmp := heap.arr[i]
	heap.arr[i] = heap.arr[j]
	heap.arr[j] = tmp
}

func main2() {
	smallHeap := NewSmallHeap(5)
	fmt.Println(smallHeap.count, smallHeap.num, len(smallHeap.arr))
	origin := []*Pair{
		{ID: "0", Value: 0},
		{ID: "7", Value: 7},
		{ID: "5", Value: 5},
		{ID: "20", Value: 20},
		{ID: "4", Value: 4},
		{ID: "1", Value: 1},
		{ID: "19", Value: 19},
		{ID: "13", Value: 13},
		{ID: "8", Value: 8},
	}
	//for _, pair := range origin {
	//	smallHeap.Insert(pair)
	//}
	////
	//for i, pair := range smallHeap.arr {
	//	fmt.Println("--------1111", i, pair)
	//}
	////
	//smallHeap.RemoveMin()

	smallHeap.BuildHeap(origin)
	tmp := smallHeap.Sort()
	for i, pair := range smallHeap.arr {
		fmt.Println("--------2222", i, pair)
	}

	fmt.Println("--------------------tmp:", tmp)
	for i, t := range tmp {
		fmt.Println("--------3333", i, t)
	}

	tt := smallHeap.Sort()
	for i, t := range tt {
		fmt.Println("--------444444", i, t)
	}

}

func (heap *SmallHeap) BuildHeap(arr []*Pair) {
	// 堆化
	for _, pair := range arr {
		heap.Insert(pair)
	}
}

func (heap *SmallHeap) Sort() []*Pair {

	fmt.Println("---------------------sortNum: ", heap.sortNum, heap.sort)
	if heap.sort {
		return heap.arr[1:heap.sortNum]
	} else {
		heap.sortNum = heap.count + 1
	}
	if heap.count < 2 {
		heap.sort = true
		return heap.arr[1:heap.sortNum]
	}
	// 堆化
	//for _, pair := range arr {
	//	heap.Insert(pair)
	//}

	//if heap.sort {
	//	return heap.arr[1:num]
	//}
	k := heap.count

	for k > 1 {
		// 将堆顶元素与最后一个元素交换位置
		k--
		// 将剩下元素重新堆化
		heap.RemoveMin()
	}
	heap.sort = true
	fmt.Println("*********************count: ", heap.count, heap.sortNum)
	return heap.arr[1:heap.sortNum]
}
