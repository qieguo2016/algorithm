/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU 缓存
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (53.45%)
 * Likes:    2565
 * Dislikes: 0
 * Total Accepted:    451.8K
 * Total Submissions: 845.4K
 * Testcase Example:
 * '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
 * 。 void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value
 * ；如果不存在，则向缓存中插入该组 key-value
 * 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
 */

// @lc code=start
#include <unordered_map>
struct LinkNode {
  int key;
  int value;
  LinkNode *pre;
  LinkNode *next;
  LinkNode() : key(0), value(0), pre(nullptr), next(nullptr){};
  LinkNode(int key_, int value_)
      : key(key_), value(value_), pre(nullptr), next(nullptr){};
};

class LRUCache {
public:
  LRUCache(int capacity) : cap(capacity), size(0) {
    head = new LinkNode();
    tail = new LinkNode();
    head->next = tail;
    tail->pre = head;
  }

  int get(int key) {
    auto it = data.find(key);
    if (it == data.end()) {
      return -1;
    }
    moveToHead(it->second);
    return it->second->value;
  }

  void put(int key, int value) {
    auto it = data.find(key);
    if (it != data.end()) {
      // update
      it->second->value = value;
      moveToHead(it->second);
      return;
    }
    // insert
    LinkNode *node = new LinkNode(key, value);
    data.emplace(key, node);
    addToHead(node);
    if (size < cap) {
      size++;
      return;
    }
    adjust();
  }

private:
  int cap;
  int size;
  std::unordered_map<int, LinkNode *> data;
  LinkNode *head;
  LinkNode *tail;

  void removeNode(LinkNode *node) {
    node->pre->next = node->next;
    node->next->pre = node->pre;
  }

  void moveToHead(LinkNode *node) {
    if (size < 2) {
      return;
    }
    removeNode(node);
    addToHead(node);
  };

  void addToHead(LinkNode *node) {
    node->pre = head;
    node->next = head->next;
    head->next->pre = node;
    head->next = node;
  }

  void adjust() {
    // clean oldest node
    auto last = tail->pre;
    removeNode(last);
    data.erase(last->key);
    delete last;
  };
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end
