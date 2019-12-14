package skiplist

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	testcase := [][]int{}
	testcaseStr := `[
		[10,10],
		[9,9],
		[10,100],
		[15,15],
		[8,8],
		[3,3],
		[8,80],
		[21,21],
		[25,25],
		[32,32]
	]`

	json.Unmarshal([]byte(testcaseStr), &testcase)

	s := NewSkipList()
	for i, data := range testcase {
		s.Insert(data[0], data[1])
		fmt.Printf("test[%d]: skiplist insert, index=%d, data=%+v\n", i, data[0], data[1])
	}
	fmt.Println("====== delete =======")
	s.Delete(9)
	s.Delete(21)
	fmt.Println("====== search =======")
	for i, data := range testcase {
		index := data[0]
		node := s.Search(index)
		fmt.Printf("test[%d]: skiplist search, index=%d, data=%+v\n", i, index, node)
	}
	fmt.Printf("skiplist rank (2, 3) = %+v\n", s.Rank(2, 3))
}
