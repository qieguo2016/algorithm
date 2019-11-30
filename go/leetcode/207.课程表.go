/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (48.34%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 32.7K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 n 门课需要选，记为 0 到 n-1。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
 *
 * 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 *
 * 示例 2:
 *
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 *
 * 说明:
 *
 *
 * 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 *
 *
 * 提示:
 *
 *
 * 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
 * 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
 *
 * 拓扑排序也可以通过 BFS 完成。
 *
 *
 *
 */
// package leetcode

// 有向图判断存在环，根据每个节点的入度数来判断，从起始点开始（入度为0）遍历，遍历时将当前节点的入度-1，
// 不断从入度为0的节点开始遍历，最后剩下节点入度>0的节点就是出现环的节点了
// 采用广度遍历，用一个队列保存将要遍历的节点

// @lc code=start
import (
	"container/list"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)  // 节点入度数组
	// 构造图，采用邻接表示方式
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, req := range prerequisites {
		// 连接方式 1 -> 0
		graph[req[1]] = append(graph[req[1]], req[0])
		in[req[0]]++
	}
	q := list.New()  // 保存入度为0的点（起始点）
	for i := range in {
		if in[i] == 0 {
			q.PushBack(i)
		}
	}
	for q.Len() > 0 {
		cur := q.Back()
		q.Remove(cur)
		val := cur.Value.(int)
		// 遍历i的邻接点
		for _, el := range graph[val] {
			in[el]--
			if in[el] == 0 {
				q.PushBack(el)  // 新的起始点
			}
		}
	}
	for _, el := range in {
		if el != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
