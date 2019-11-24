// package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (61.14%)
 * Likes:    129
 * Dislikes: 0
 * Total Accepted:    13.4K
 * Total Submissions: 21.8K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
 */

// @lc code=start
type Node struct {
	Child map[rune]*Node
	Value bool
}

type Trie struct {
	Root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Root: &Node{
			Child: map[rune]*Node{},
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.Root
	for _, c := range word {
		if _, ok := cur.Child[c]; !ok {
			cur.Child[c] = &Node{
				Child: map[rune]*Node{},
			}
		}
		cur = cur.Child[c]
	}
	cur.Value = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.Root
	for _, c := range word {
		if _, ok := cur.Child[c]; !ok {
			return false
		}
		cur = cur.Child[c]
	}
	return cur.Value
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.Root
	for _, c := range prefix {
		if _, ok := cur.Child[c]; !ok {
			return false
		}
		cur = cur.Child[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
