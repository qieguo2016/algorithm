package base

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

/**
 * 冒泡排序
 * @param       : <Array> target数组
 * @description : 冒泡排序，更贴切的形容应该是沉底排序，每一轮内循环就是最大数沉底了。
 */
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

/**
 * 选择排序
 * @param       : <Array> target数组
 * @description : 一次内循环得到最大值的下标，然后只交换一次次序，将最大值和内循环末尾对调。
 */
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


