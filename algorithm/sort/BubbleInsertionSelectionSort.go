package main

import "fmt"

/**

 */

func main() {
	var needSortSlice []int = []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//sortedData := BubbleSort(needSortSlice)
	//sortedData := InsertionSort(needSortSlice)
	//sortedData := SelectionSort(needSortSlice)
	sortedData := ShellSort(needSortSlice)

	fmt.Println(sortedData)

}

//---------------------------------冒泡---------------------------
// 冒泡排序: 从小到大排序
// 原地排序、稳定的排序
// 时间复杂度：O(N^2)
// 冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
func BubbleSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 外层控制循环次数,外层每循环一次获取到未排序部分的最大值
	for i := 0; i < n; i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				//temp := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = temp
				arr[j], arr[j+1] = arr[j+1], arr[j]

				flag = true // 表示有数据交换
			}
		}
		// 没有数据交换，提前退出
		if !flag {
			break
		}
	}
	return arr
}

// 该冒泡的方式是每次获取最小的
func BubbleSort1(arr []int) []int {
	// 外层每循环一次获取未排序最小的值
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//---------------------------------冒泡 end---------------------------

//---------------------------------插入---------------------------
// 插入排序: 原地排序、稳定的排序
// 从尾到头比较已排序的数据
// 时间复杂度：O(N^2)
// 插入排序算法的核心思想 将数据分为：已排序区间、未排序区间。 最初，已排序区间为第一个元素，以此获取未排序区间中的数与已排序区间进行比较，找到正确的位置进行插入，并保证已排序区间的数据一直有序。直到未排序区间中元素为空
// 对于不同的查找插入点方法（从头到尾、从尾到头），元素的比较次数是有区别的。但对于一个给定的初始序列，移动操作的次数总是固定的，就等于逆序度
func InsertionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 外层每次循环都是取未排序数据中的第一个数据进行处理
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1 // 有序部分的最后一个值
		// 查找插入的位置
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j] // 数据移动
			} else {
				break
			}
		}
		arr[j+1] = value // 插入位置
	}

	return arr
}
func InsertionSort1(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 外层每次循环都是取未排序数据中的第一个数据进行处理
	for i := 1; i < n; i++ {
		//此处j代表有序区间最多循环处理的次数
		for j := 0; j < i; j++ {
			if arr[i-j] < arr[i-j-1] {
				arr[i-j], arr[i-j-1] = arr[i-j-1], arr[i-j]
			} else {
				break
			}
		}
	}

	return arr
}

//---------------------------------插入 end---------------------------

//---------------------------------选择排序---------------------------
// 选择排序: 原地排序，不稳定的排序
// 时间复杂度：O(N^2)
// 选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
func SelectionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 外层控制寻找最小值的次数，每次循环，获取未排序数据的最小值
	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

//---------------------------------插入 end---------------------------

// 希尔排序
// 希尔排序特别适合近乎有序的序列排序使用
// 先将整个待排记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录“基本有序”时，在对全体进行一次直接插入排序。
// http://www.haodaquan.com/143
// 定义说的严谨，可是吧，一时半会很难看懂。我们知道插入排序是每次都是取基准元素左侧的元素比较大小，也就是每次比较的步长为1，如果是近乎有序的数组，插入排序很容易退化成O（n^2），而希尔排序每次选择的步长是不一样的（当然最后的步长肯定为1，也就是最后用一次插入排序），这是打破插入排序O（n^2）时间复杂度的关键步骤。
func ShellSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	var i, j, gap int
	// gap 代表分成了多少组
	for gap = n / 2; gap > 0; gap = gap / 2 { //增量起始值为n/2,之后逐次减半
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i = gap; i < n; i++ {
			//i:代表即将插入的元素角标，作为每一组比较数据的最后一个元素角标
			//j:代表与i同一组的数组元素角标
			for j = i - gap; j >= 0 && arr[j] > arr[j+gap]; j = j - gap { ////在此处 - gap 为了避免下面数组角标越界
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}

	return arr
}
