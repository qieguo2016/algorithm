package main

type countNode struct {
	start      int
	leftCount  int
	rightCount int
}

// 字符串分隔计数
// 给定字符串，**|***|*|，统计给定区间组内各自被||围起来的**数量
// @s **|***|*|**，注意左右两侧无|线围起来不算
// @indices, [[1,5], [2,6]]，表示统计[1,5]内被围起来的**数量，双闭合，字符串从1开始计数
// 解法如下：
//        * * | * * * | * | * * | *
// left   0 0 0 0 0 0 3 3 4 4 4 6 6
// right  0 0 3 3 3 3 4 4 6 6 6 0 0
// 计数比较简单，末尾向左取整，起始部分如果是|则算上本段，否则算到下一段
func stringSplitCount(s string, indices [][]int) []int {
	// ret := []int{}
	// splitPos := make([]int, 0)
	// for i, c := range s {
	// 	if c == '|' {
	// 		splitPos = append(splitPos, i)
	// 	}
	// }
	// // 二分法找第一个比startIdx大的位置，再找第一个比endIdx小的位置
	// for i := 0; i < len(indices); i++ {
	// 	cur := indices[i]
	// 	a := searchBigger(splitPos, cur[0]-1)
	// 	b := searchSmaller(splitPos, cur[0]-1)
	// 	ret = append(ret, splitPos[b]-splitPos[a]-b+a) // note: 有边界问题
	// }

	ret := []int{}
	counts := make([]*countNode, len(s))
	// 遍历s，|*** 为一段（开头**自成一段），段内公用一套leftCount、rightCount、start
	// 到达|时，累计计算本段的leftCount，利用指针来更新上一段的rightCount
	for i, c := range s {
		if c == '*' {
			if i <= 0 {
				counts[i] = &countNode{
					start:      -1, // 标识开头的*
					leftCount:  0,
					rightCount: 0,
				}
			} else {
				counts[i] = counts[i-1] // 利用指针串联整个区间
			}
		} else if c == '|' {
			cur := &countNode{
				start:     i,
				leftCount: 0,
			}
			if i > 0 {
				pre := counts[i-1] // 利用指针改掉整个区间
				if pre.start >= 0 {
					cur.leftCount = i - pre.start - 1 + pre.leftCount
					pre.rightCount = cur.leftCount
				}
			}
			counts[i] = cur
		}
	}

	//        * * | * * * | * | * * | *
	// left   0 0 0 0 0 0 3 3 4 4 4 6 6
	// right  0 0 3 3 3 3 4 4 6 6 6 0 0
	// 计数比较简单，末尾向左取整，起始部分如果是|则算上本段，否则算到下一段
	for i := 0; i < len(indices); i++ {
		index := indices[i]
		end := counts[index[1]-1].leftCount
		c := counts[index[0]-1]
		start := c.rightCount
		if s[index[0]-1] == '|' {
			start = c.leftCount
		}
		ret = append(ret, end-start)
	}

	return ret
}
