package main

// 购物车抽奖
func food(code [][]int, cart []int) bool {
	return (&foodSol{}).call(code, 0, cart, 0)
}

type foodSol struct{}

func (s *foodSol) call(code [][]int, cs int, food []int, fs int) bool {
	// i == 0
	if cs >= len(code) {
		return true
	}

	if fs >= len(food) {
		return false
	}

	// [[1,2] [2,-1, 3]]    [2,1,2,2,3,3]
	for i := fs; i < len(food); i++ {
		match := true
		j := 0
		for j < len(code[cs]) {
			if code[cs][j] != -1 && code[cs][j] != food[i+j] {
				match = false
				break
			}
			j++
		}
		if match && s.call(code, cs+1, food, i+j) {
			return true
		}

	}
	return false
}
