package main

/**
（1）O(n) 时间复杂度内求无序数组中的第 K 大元素？
如果 p+1=K，那 A[p] 就是要求解的元素（如果每次分区都很均衡的话，每次都需要遍历剩余数量的1/2）
*/

//---------------------------------归并排序---------------------------
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

// 非原地排序、稳定的排序
// 最好情况、最坏情况，还是平均情况，时间复杂度都是：O(NlogN)
// 空间复杂度： O(N)
// 核心思想还是蛮简单的。如果要排序一个数组，我们先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
func mergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	// 将剩余的数据拷贝到临时数组tmp
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmpArr)
}

// ---------------------------------快排排序---------------------------
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// 原地排序、非稳定的排序
// 大部分情况下的时间复杂度都可以做到 O(nlogn)，只有在极端情况下，才会退化到 O(n^2)
// 空间复杂度： O(1)
// 快排的思想是这样的：如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
//我们遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的。
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	// 分区函数实际上我们前面已经讲过了，就是随机选择一个元素作为 pivot（一般情况下，可以选择 p 到 r 区间的最后一个元素），然后对 A[p…r] 分区，函数返回 pivot 的下标。
	q := partition(arr, start, end) // 获取分区点
	quickSort(arr, start, q-1)
	quickSort(arr, q+1, end)
}

func partition(arr []int, start, end int) int {
	//初时 选取最后一位当对比数字
	pivot := arr[end]
	var i = start // 通过下标 i 将arr分成两部分，arr[start...i-1]的元素都是pivot的值，我们可以称为"已处理区间"
	//j用来控制遍历的数据的次数
	for j := start; j <= end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
