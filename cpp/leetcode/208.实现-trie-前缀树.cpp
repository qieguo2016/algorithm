/*
 * @lc app=leetcode.cn id=208 lang=cpp
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode.cn/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (71.92%)
 * Likes:    1454
 * Dislikes: 0
 * Total Accepted:    263.1K
 * Total Submissions: 365.7K
 * Testcase Example:
 * '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Trie（发音类似 "try"）或者说 前缀树
 * 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
 *
 * 请你实现 Trie 类：
 *
 *
 * Trie() 初始化前缀树对象。
 * void insert(String word) 向前缀树中插入字符串 word 。
 * boolean search(String word) 如果字符串 word 在前缀树中，返回
 * true（即，在检索之前已经插入）；否则，返回 false 。 boolean startsWith(String
 * prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true
 * ；否则，返回 false 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
 * [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
 * 输出
 * [null, null, true, false, true, null, true]
 *
 * 解释
 * Trie trie = new Trie();
 * trie.insert("apple");
 * trie.search("apple");   // 返回 True
 * trie.search("app");     // 返回 False
 * trie.startsWith("app"); // 返回 True
 * trie.insert("app");
 * trie.search("app");     // 返回 True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * word 和 prefix 仅由小写英文字母组成
 * insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次
 *
 *
 */
#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Trie {
public:
  Trie() : children_(26), isEnd_(false) {}

  void insert(string word) {
    auto node = this;
    for (auto c : word) {
      c -= 'a';
      if (node->children_[c] == nullptr) {
        node->children_[c] = new Trie();
      }
      node = node->children_[c];
    }
    node->isEnd_ = true;
  }

  bool search(string word) {
    auto node = searchPrefix(word);
    return node != nullptr && node->isEnd_;
  }

  bool startsWith(string prefix) { return searchPrefix(prefix) != nullptr; }

private:
  vector<Trie *> children_;
  bool isEnd_;

  Trie *searchPrefix(string& prefix) {
    auto node = this;
    for (auto c : prefix) {
      c -= 'a';
      if (node->children_[c] == nullptr) {
        return nullptr;
      }
      node = node->children_[c];
    }
    return node;
  }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
// @lc code=end
