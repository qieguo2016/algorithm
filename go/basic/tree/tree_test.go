package sort

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBPTree(t *testing.T) {
	order := 5
	bp := NewBPTree(order)
	fmt.Printf("NewBPTree, order = %d\n", order)

	input := [][]int{}
	inputStr := `[
		[5,5],
		[8,8],
		[10,10],
		[15,15],
		[16,16],

		[17,17],
		[18,18],
		[6,6],
		[9,9],

		[19,19],

		[20,20],
		[21,21],
		[22,22],

		[25,25],
		[26,26],
		[27,27],
		[28,28],
		[29,29],

		[29],
		[27],
		[26],
		[25],

		[251],
		[0],
		[1],
		[2],

		[5],
		[10]

	]`
	json.Unmarshal([]byte(inputStr), &input)

	output := []interface{}{}
	outStr := "[null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]"
	json.Unmarshal([]byte(outStr), &output)

	for i, data := range input {
		if len(data) > 1 {
			fmt.Printf("insert dataset %d, data=%+v\n", i, data)
			bp.Insert(data[0], data[1])
		} else {
			res := bp.Search(data[0])
			fmt.Printf("search dataset %d, index=%d, res=%+v\n", i, data[0], res)
			res = bp.Delete(data[0])
			fmt.Printf("delete dataset %d, index=%d, res=%+v\n", i, data[0], res)
			res = bp.Search(data[0])
			fmt.Printf("search dataset %d, index=%d, res=%+v\n", i, data[0], res)
		}
	}
	root := bp.RootPage
	for root != nil {
		fmt.Printf("root = %+v\n", root)
		root = root.Next
	}
	leaf := bp.LeafPage
	for leaf != nil {
		fmt.Printf("leaf = %+v\n", leaf)
		leaf = leaf.Next
	}
}
