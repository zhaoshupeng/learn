package main

import (
	"math"
)

/**
	能做到线性的时间复杂度，主要原因是，这三个算法是非基于比较的排序算法，都不涉及元素之间的比较操作。
（1）线性排序：如何根据年龄给 100 万用户排序？
*/

//----------------------------桶排序-------------------------------------------------
// ---第一种

func getMinAndMax(arr []int) (int, int) {
	// 数组最小值
	minValue := arr[0]
	// 数组最大值
	maxValue := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		} else if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}

	return minValue, maxValue
}

/**
// 非原地排序、稳定的排序
// 时间复杂度：整个桶排序的时间复杂度就是 O(n*log(n/m))。当桶的个数 m 接近数据个数 n 时，log(n/m) 就是一个非常小的常量，这个时候桶排序的时间复杂度接近 O(n)。
// 空间复杂度： O(N)
// 核心思想是将要排序的数据分到几个有序的桶里，每个桶里的数据再单独进行排序。桶内排完序之后，再把每个桶里的数据按照顺序依次取出，组成的序列就是有序的了。

适用场景：桶排序比较适合用在外部排序中。所谓的外部排序就是数据存储在外部磁盘中，数据量比较大，内存有限，无法将数据全部加载到内存中。
使用限制：
	（1）首先，要排序的数据需要很容易就能划分成 m 个桶，并且，桶与桶之间有着天然的大小顺序。
	（2）其次，数据在各个桶之间的分布是比较均匀的。在极端情况下，如果数据都被划分到一个桶里，那就退化为 O(nlogn) 的排序算法了。
*/

//BucketSort : arr 待排序数组，bucketSize表示每个桶的取值区间跨度
func BucketSort(arr []int, bucketSize int) {
	num := len(arr)
	if num <= 1 {
		return
	}
	minValue, maxValue := getMinAndMax(arr)
	bucketCount := (maxValue-minValue)/bucketSize + 1 // 注意需要加1
	buckets := make([][]int, bucketCount)             // 二维切片
	//// 二维数组初始化
	//for i := range buckets {
	//	buckets[i] = make([]int, 0)
	//}
	// 将数组中值分配到各个桶里
	for i := 0; i < num; i++ {
		bucketIndex := (arr[i] - minValue) / bucketSize             // 桶序号，第一个桶的最小值是minValue
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i]) // 加入对应的桶中
	}

	// 排序
	k := 0
	for i := 0; i < bucketCount; i++ {
		length := len(buckets[i])
		if length == 0 {
			continue
		}
		QuickSort(buckets[i]) //快排
		for j := 0; j < length; j++ {
			arr[k] = buckets[i][j]
			k++
		}
	}

}

//-------------第二种------------

func BucketSortSecond(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	max := getMax(arr)

	buckets := make([][]int, length) // 二维切片,桶的个数是length

	// 将数据放到不同的桶中
	index := 0
	for i := 0; i < length; i++ {
		index = arr[i] * (length - 1) / max             // 桶序号, (length - 1) 最大值分到的桶的序号，可以理解为arr[i]的值越接近max，桶的位置越接近最后一个桶
		buckets[index] = append(buckets[index], arr[i]) // 加入对应的桶中
	}
	//fmt.Println("分桶后的结果----", buckets)

	tmpPos := 0 // 标记数组位置
	for i := 0; i < length; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i]) // 桶内做快速排序
			//fmt.Println("桶内排序后--", buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += len(buckets[i])
		}

	}

}

// 获取待排序数组中的最大值
func getMax(arr []int) int {
	max := arr[0]
	// range在遍历大数据时候，效率不高，需要拷贝数据: range使用的是副本，也就是键和值都是定义了一个新的变量，也就是值拷贝，
	//for _, val := range arr {
	//	if max < val {
	//		max = val
	//	}
	//}
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

//-------------------------计数排序--------------------------------------------------------------------

/**
问题：如果你所在的省有 50 万考生，如何通过成绩快速排序得出名次呢？
*/
/**
使用场景和限制：
	(1) 计数排序只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大很多，就不适合用排序了。
	(2) 排序只能给非负整数排序，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数。
*/
// 时间复杂度O(N)
// 计数排序其实是桶排序的一种特殊情况。当要排序的 n 个数据，所处的范围并不大的时候，比如最大值是 k，我们就可以把数据划分成 k 个桶。每个桶内的数据值都是相同的，省掉了桶内排序的时间。
// 假设数组中存储的都是非负整数

func CountingSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	// 查找数组中数据的范围
	var max int = math.MinInt32
	//var max int = arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 这种利用另外一个数组来计数的实现方式是不是很巧妙呢？这也是为什么这种排序算法叫计数排序的原因。
	c := make([]int, max+1) // 申请一个计数数组c，下标大小是[0,max],其中的值是对应的值的个数

	// 计算每个元素的个数，放入 c 中
	for i := range arr {
		c[arr[i]]++
	}
	// 依次累加，然后c[k]里存储的就是小于等于k的个数
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组 r，存储排序之后的结果
	r := make([]int, length)
	// 从后向前依次扫描数组arr(这样做数为了保证算法的稳定性)
	for i := length - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}

	// 将结果拷贝给 a 数组
	copy(arr, r)
}

//-----------------------------------基数排序（Radix sort）-----------------------------------------------------------------------------------------
/**
问题：
	(1) 假设我们有 10 万个手机号码，希望将这 10 万个手机号码从小到大排序，你有什么比较快速的排序方法呢？
         桶排序、计数排序能派上用场吗？手机号码有 11 位，范围太大，显然不适合用这两种排序算法。
*/
//RadixSort 时间复杂度O(N)
// 核心思想：借助稳定排序算法，先按照最后一位来排序，然后，再按照倒数第二位重新排序，以此类推，最后按照第一位重新排序。（单词这类位数不够的可以在后面补“0”）
//根据每一位来排序，我们可以用刚讲过的桶排序或者计数排序，它们的时间复杂度可以做到 O(n)。如果要排序的数据有 k 位，那我们就需要 k 次桶排序或者计数排序，总的时间复杂度是 O(k*n)。

/**
使用场景和限制：
	（1）基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系
	（2）基数排序算法需要借助桶排序或者计数排序来完成每一个位的排序工作。
*/

// arr := []int{3,38,11,659,162}

func RadixSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	max := getMax(arr)

	// 从个位开始，对数组arr按"指数"进行排序
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortForRadix(arr, exp)
	}
}

// 计数排序-对数组按照"某个位数"进行排序 exp 指数
func countingSortForRadix(arr []int, exp int) {
	// 主题逻辑还是计数排序
	length := len(arr)
	if length <= 1 {
		return
	}

	// 统计长度 0~9(这是每一位的取值范围)
	c := make([]int, 10)

	for i := 0; i < length; i++ {
		// 余数，即对应的位上的值
		num := (arr[i] / exp) % 10

		// 计算每个元素的个数; 即统计余数相等的个数，进行递增
		c[num]++
	}

	// 目的是增加稳定性
	// 依次累加，然后c[k]里存储的就是小于等于k的个数
	for i := 1; i < len(c); i++ {
		//c[i] = c[i-1] + c[i]
		c[i] += c[i-1]
	}

	// 临时数组 r，存储排序之后的结果
	r := make([]int, length)

	// 从后向前依次扫描数组arr(这样做数为了保证算法的稳定性)
	for i := length - 1; i >= 0; i-- {
		// 余数
		num := (arr[i] / exp) % 10
		index := c[num] - 1

		r[index] = arr[i]
		// 统计余数相等的个数，进行递减
		c[num]--
	}

	// 将结果拷贝给 a 数组
	copy(arr, r)
}
