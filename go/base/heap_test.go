package base

import (
	"fmt"
	"testing"
)

func printHeap(arr []int) {
	length := len(arr)
	queue := []int{0}
	n := 0
	for n < length {
		for _, c := range queue[n:len(queue)] {
			fmt.Printf("%d ", arr[c])
			n++
			if 2*c+1 < length {
				queue = append(queue, 2*c+1)
			}
			if 2*c+2 < length {
				queue = append(queue, 2*c+2)
			}
		}
		fmt.Print("\n")
	}
}

func TestSmallRootHeap(t *testing.T) {
	for i := 0; i < 2; i++ {
		arr := makeRandomArray(7)
		fmt.Printf("origin arr=%v\n", arr)
		NewSmallRootHeap(arr)
		fmt.Printf("heap arr=%v\n", arr)
		printHeap(arr)
	}
}
