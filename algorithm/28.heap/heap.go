package main

/**
问题：
（1）在实际的软件开发中，快速排序的性能要比堆排序好，这是为什么呢？
*/
/**
只要满足这两点，它就是一个堆:
	(1)堆是一个完全二叉树(完全二叉树要求，除了最后一层，其他层的节点个数都是满的，最后一层的节点都靠左排列。完全二叉树适合用数组存储，
数组中下标为 i的节点的左子节点，就是下标为i∗2 的节点，右子节点就是下标i*2+1 的节点，父节点就是下标为 i/2的节点。)；
	(2)堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。
*/

func main() {

}

// 以大顶堆 为例
type Heap struct {
	arr   []int //存储堆的数组，从下标 1 开始存储数据
	num   int   // 堆可以存储的最大数据个数
	count int   // 堆中已经存储的数据个数
}

//init heap
//初始化堆/建堆，capacity
func NewHeap(capacity int) *Heap {
	heap := Heap{}
	heap.arr = make([]int, capacity+1) // 因为从下标1 开始存储数据
	heap.num = capacity
	heap.count = 0
	return &heap
}

//top-max heap -> heapify from down to up
// 新增一个数据
func (heap *Heap) Insert(data int) {
	//defensive
	if heap.count == heap.num { // 堆满了
		return
	}
	heap.count++
	heap.arr[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.arr[i] > heap.arr[parent] { // 自下往上堆化
		swap(heap.arr, i, parent)
		i = parent
		parent = i / 2
	}
}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
