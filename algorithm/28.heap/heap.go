package main

import "fmt"

/**
问题：
（1）在实际的软件开发中，快速排序的性能要比堆排序好，这是为什么呢？
	第一点，堆排序数据访问的方式没有快速排序友好。对于快速排序来说，数据是顺序访问的。而对于堆排序来说，数据是跳着访问的。
	第二点，对于同样的数据，在排序过程中，堆排序算法的数据交换次数要多于快速排序。
（2）如何基于堆实现排序？
	借助堆这种数据结构实现的排序就是堆排序，这种排序方法时间复杂度非常稳定，是O(NlogN),堆排序的过程分为：建堆和排序
	1）建堆
     a) 方式一：借助在堆中插入一个元素堆思路，起初堆中只包含一个数据，就是下标1，然后一次插入剩余数据。第一种建堆思路的处理过程是从前往后处理数组数据，并且每个数据插入堆中时，都是从下往上堆化。
	 b) 方式二：实现思路，是从后往前处理数组，并且每个数据都是从上往下堆化。因为叶子节点往下堆化只能自己跟自己比较，所以我们直接从第一个非叶子节点开始，依次堆化就行了。
*/
/**
只要满足这两点，它就是一个堆:
	(1)堆是一个完全二叉树(完全二叉树要求，除了最后一层，其他层的节点个数都是满的，最后一层的节点都靠左排列。完全二叉树适合用数组存储，
数组中下标为 i的节点的左子节点，就是下标为i∗2 的节点，右子节点就是下标i*2+1 的节点，父节点就是下标为 i/2的节点。)；
	(2)堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。
*/

/**
适用场景：
（1）优先级队列
	1）合并有序小文件
		我们将从小文件中取出来的字符串放入到小顶堆中，那堆顶的元素，也就是优先级队列队首的元素，就是最小的字符串。我们将这个字符串放入到大文件中，并将其从堆中删除。
	2) 高性能定时器
（2）利用堆求 Top K
	我们可以维护一个大小为 K 的小顶堆，顺序遍历数组，从数组中取出数据与堆顶元素比较。如果比堆顶元素大，我们就把堆顶元素删除，并且将这个元素插入到堆中；

（3）利用堆求中位数
	我们需要维护两个堆，一个大顶堆，一个小顶堆。大顶堆中存储前半部分数据，小顶堆中存储后半部分数据，且小顶堆中的数据都大于大顶堆中的数据。
如果新加入的数据小于等于大顶堆的堆顶元素，我们就将这个新数据插入到大顶堆；否则，我们就将这个新数据插入到小顶堆。（然后调整两个堆中的数量，让大顶堆的堆顶元素是中位数）
*/

func main1() {
	arr := []int{0, 7, 5, 20, 4, 1, 19, 13, 8}
	Sort(arr, len(arr)-1)
	fmt.Println("after sort", arr)
}

// 以大顶堆 为例
type Heap struct {
	arr   []int //存储堆的数组，从下标 1 开始存储数据
	num   int   // 堆可以存储的最大数据个数
	count int   // 堆中已经存储的数据个数
}

// init heap
// 初始化堆/建堆，capacity
func NewHeap(capacity int) *Heap {
	heap := Heap{}
	heap.arr = make([]int, capacity+1) // 因为从下标1 开始存储数据
	heap.num = capacity
	heap.count = 0
	return &heap
}

/**
一个包含n个节点的完全二叉树，其树高不超过log2N，堆化的时间复杂度跟树高成正比，也就是O(logN),插入和删除堆顶元素的主要逻辑就是堆化
*/

// top-max heap -> heapify from down to up
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
		swap(heap.arr, i, parent) //交换数组中的两个值
		i = parent
		parent = i / 2
	}
}

func (heap *Heap) RemoveMax() {
	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	//删除堆顶元素，并将堆的最后的值放置到堆顶，然后从堆顶进行堆化，这样结束后还是能满足完全二叉树的特性(或堆的定义)
	swap(heap.arr, 1, heap.count)
	heap.count--

	//heapify from up to down
	// 从上向下堆化
	heapifyUpToDown(heap.arr, heap.count)
}

// swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// heapify
func heapifyUpToDown(arr []int, count int) { // 自上往下堆化
	for i := 1; i <= count/2; {
		maxIndex := i

		if arr[i] < arr[i*2] {
			maxIndex = arr[i*2]
		}
		if i*2+1 <= count && arr[maxIndex] < arr[i*2+1] {
			maxIndex = arr[i*2+1]
		}
		// 已经大于左右子树的任意值
		if maxIndex == i {
			break
		}
		swap(arr, i, maxIndex)
		i = maxIndex
	}
}

// 堆排序是一种原地的、时间复杂度为 O(nlogn) 的排序算法。
// 建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，
// 也就是最大的元素。我们把它跟最后一个元素交换，那最大元素就放到了下标为 n 的位置。(类似于删除堆顶元素，)
// n 表示数据的个数，数组 a 中的数据从下标 1 到 n 的位置。
func Sort(arr []int, n int) {
	BuildHeap(arr, n)
	k := n

	for k > 1 {
		// 将堆顶元素（最大）与最后一个元素交换位置
		swap(arr, 1, k)
		k--
		// 将剩下元素重新堆化，堆顶元素变成最大元素
		heapify(arr, k, 1)
	}

}

// ---第二种方式建堆
// 建堆时间复杂度是O(N): 堆化的节点从倒数第二层开始。每个节点堆化的过程中，需要比较和交换的节点个数，跟这个节点的高度 K 成正比。我们只需要将每个节点的高度求和，得出的就是建堆的时间复杂度。
func BuildHeap(arr []int, n int) {
	// n / 2 为最后一个叶子节点的父节点
	// 也就是最后一个非叶子节点，依次堆化直到根节点
	for i := n / 2; i >= 1; {
		// 对下标从n/2开始到1的数据进行堆化，下标是n/2+1到n的节点是叶子结点，不需要进行堆化
		heapify(arr, n, i)
		i -= 1

	}
}

// a 要堆化的数组,n 最后堆元素下标,i 要堆化的元素下标
func heapify(a []int, n int, i int) {
	//maxPos := i
	for {

		maxPos := i // 最大值位置
		// 与左子节点比较，获取最大值位置
		if i*2 <= n && a[i] < a[i*2] {
			maxPos = i * 2
		}

		// 最大值与右子节点比较，获取最大值位置
		if i*2+1 <= n && a[maxPos] < a[i*2+1] {
			maxPos = i*2 + 1
		}

		// 最大值是当前位置结束循环
		if maxPos == i {
			break
		}
		// 与子节点交换位置
		swap(a, i, maxPos)
		// 以交换后子节点位置接着往下查找
		i = maxPos
	}
}
