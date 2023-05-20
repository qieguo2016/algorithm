/*
 * @lc app=leetcode.cn id=211 lang=cpp
 *
 * [211] 添加与搜索单词 - 数据结构设计
 *
 * https://leetcode.cn/problems/design-add-and-search-words-data-structure/description/
 *
 * algorithms
 * Medium (49.62%)
 * Likes:    497
 * Dislikes: 0
 * Total Accepted:    72.3K
 * Total Submissions: 145.6K
 * Testcase Example:
 * '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * 请你设计一个数据结构，支持 添加新单词 和
 * 查找字符串是否与任何先前添加的字符串匹配 。
 *
 * 实现词典类 WordDictionary ：
 *
 *
 * WordDictionary() 初始化词典对象
 * void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
 * bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true
 * ；否则，返回  false 。word 中可能包含一些
 * '.' ，每个 . 都可以表示任何一个字母。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 *
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * 输出：
 * [null,null,null,null,false,true,true,true]
 *
 * 解释：
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // 返回 False
 * wordDictionary.search("bad"); // 返回 True
 * wordDictionary.search(".ad"); // 返回 True
 * wordDictionary.search("b.."); // 返回 True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= word.length <= 25
 * addWord 中的 word 由小写英文字母组成
 * search 中的 word 由 '.' 或小写英文字母组成
 * 最多调用 10^4 次 addWord 和 search
 *
 *
 */

#include <string>
#include <vector>
using namespace std;

// @lc code=start
// 用前缀树解法，当匹配到.时往下搜索所有子节点
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

  bool search(string word) { return search(this, word, 0); }

private:
  vector<Trie *> children_;
  bool isEnd_;

  bool search(Trie *node, string &word, int idx) {
    if (idx == word.size()) {
      return node->isEnd_;
    }
    char c = word[idx];
    if (c == '.') {
      for (auto child : node->children_) {
        if (child != nullptr && search(child, word, idx + 1)) {
          return true;
        }
      }
    } else {
      c -= 'a';
      if (node->children_[c] != nullptr &&
          search(node->children_[c], word, idx + 1)) {
        return true;
      }
    }
    return false;
  }
};

class WordDictionary {
private:
  Trie *root_;

public:
  WordDictionary() { root_ = new Trie(); }

  void addWord(string word) { root_->insert(word); }

  bool search(string word) { return root_->search(word); }
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */
// @lc code=end
