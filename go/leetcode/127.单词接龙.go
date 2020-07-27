/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (36.59%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 39.1K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 * 
 * 
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 * 
 * 
 * 说明:
 * 
 * 
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * 
 * 输出: 5
 * 
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 * 
 * 输出: 0
 * 
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 * 
 */

// @lc code=start


func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord || len(wordList) <= 0 {
		return 0
	}
	wordMap := map[string][]string{}
	isEndWordInList := false
	for _, w := range wordList {
		if w == beginWord {
			continue // 避免形成环
		}
		if w == endWord {
			isEndWordInList = true // 最后一个单词要在单词列表中
		}
		for i := range w {
			k := w[:i] + "*"
			if i < len(w)-1 {
				k += w[i+1:] // 用*it/h*t/hi* 的形式归一化，不需要遍历26个字母
			}
			if _, ok := wordMap[k]; ok {
				wordMap[k] = append(wordMap[k], w)
			} else {
				wordMap[k] = []string{w}
			}
		}
	}
	if !isEndWordInList {
		return 0
	}
	visited := map[string]bool{}
	queue := []string{beginWord}
	depth := 1
	// 广度（层序）遍历得到的结果必要是层数最少的
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			w := queue[i]
			if w == endWord {
				return depth
			}
			if visited[w] {
				continue
			}
			for j := range w {
				k := w[:j] + "*"
				if j < len(w)-1 {
					k += w[j+1:] // 用*it/h*t/hi* 的形式归一化，不需要遍历26个字母
				}
				if _, ok := wordMap[k]; ok {
					queue = append(queue, wordMap[k]...)
				}
			}
			visited[w] = true
		}
		queue = queue[size:]
		depth++
	}
	return 0
}

// @lc code=end

