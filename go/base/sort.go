package base

// BubbleSort 冒泡排序
// 冒泡排序，更贴切的形容应该是沉底排序，每一轮内循环就是最大数沉底了。
func BubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for j := len(arr); j > 0; j-- {
		for i := 1; i < j; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
	return arr
}

// SelectSort 选择排序
// 一次内循环得到最大值的下标，然后只交换一次次序，将最大值和内循环末尾对调。
func SelectSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for j := len(arr); j > 0; j-- {
		maxIndex := 0
		for i := 1; i < j; i++ {
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		swap(arr, maxIndex, j-1)
	}
	return arr
}

// QuickSort 非原地快排
// 取第一个元素为基准将数组分为左右两个数组，然后分别递归左右数组，最后进行合并
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := []int{}
	right := []int{}
	datum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < datum {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	left = append(left, datum)
	right = QuickSort(right)
	return append(left, right...)
}

// QuickSortInPlace 原地快排
// 非原地快排每层循环都开辟两个新数组，浪费空间。采用左右双指针往中间扫描，根据大小两两交换可以节省空间
func QuickSortInPlace(target []int) []int {
	var quickSortInPlace func(arr []int, left int, right int) []int
	quickSortInPlace = func(arr []int, left int, right int) []int {
		// 原地快排内层函数, 采用切片参数，减少值传递copy
		// slice也是传值，但是slice本身只是一个指向底层数组的指针包装
		if left >= right {
			return arr
		}
		datum := arr[left]
		i := left
		j := right
		// 以双指针相遇地将数组分割成两部分，将两侧分布不对的元素互换
		for i < j {
			// 直接找到比基准小的元素
			for i < j && arr[j] >= datum {
				j--
			}
			arr[i] = arr[j]
			// 直接找到比基准大的元素
			for i < j && arr[i] <= datum {
				i++
			}
			arr[j] = arr[i]
		}
		arr[i] = datum // i就是当前的分割线
		quickSortInPlace(arr, left, i-1)
		quickSortInPlace(arr, i+1, right)
		return arr
	}
	return quickSortInPlace(target, 0, len(target)-1)
}

//
