package base

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {
	cache := NewLFUCache(10)

	input := [][]int{}
	inputStr := "[[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]"

	json.Unmarshal([]byte(inputStr), &input)

	out := []interface{}{}
	outStr := "[null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]"
	json.Unmarshal([]byte(outStr), &out)

	for i, data := range input {
		if len(data) > 1 {
			cache.Put(data[0], data[1])
		} else {
			v := cache.Get(data[0])
			expect, ok := out[i].(float64)
			if !ok {
				fmt.Printf("expect is null, k=%d, v=%d", data[0], v)
				t.Errorf("expect is null, k=%d, v=%d", data[0], v)
			}
			if v != int(expect) {
				fmt.Printf("expect=%v, k=%d, v=%d", int(expect), data[0], v)
				t.Errorf("expect != v, k=%d, v=%d", data[0], v)
			}
		}
	}
}
