package base

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}