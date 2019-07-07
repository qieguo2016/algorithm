# 算法与数据结构

用通俗易懂的语言来介绍工作和面试中常见的数据结构和算法，提供golang和js两种语言的实现。另外提供面试中常见算法题，尤其是leetcode题目的讲解和golang代码实现。

## 数据结构部分

#### [线性表：数组与链表、队列与栈](/md/数组与链表，堆栈与队列.md)

介绍了基本数据结构之线性表的特点和原理以及js的算法实现。
由于js语言本身的特点，线性表在js中主要以数组的应用为主，而js数组本身也并不是传统意义上的连续线性表。

#### [哈希表：js对象与哈希表](/md/js对象与哈希表.md)

js对象大家都用的很多，但其底层的哈希表特性你是否清楚？
在这篇文章里，小茄会用最平白易懂的语言来讲哈希结构的原理、构造方法，当然，还有哈希表在js中的应用。

#### [树：js中的树与二叉树](/md/js中的树.md)

树是一种带有层次的数据结构，分层特性可以用于实现数据存储和快速查找。
比较常见的应用场景就是各种目录结构，如文件目录、DOM结构等，由于每指定一层就是一层筛选，所以可以用于快速查找。
js中可以通过对象的哈希结构来实现树结构，两种数据结构结合，速度更快。

#### [堆 (golang)](/go/base/heap.go)

堆是一种带有顺序结构的完全二叉树，分为大根堆和小根堆，根据完全二叉和父子大小关系，利用数组结构比较容易实现堆结果。另外golang本身的堆实现(container/heap.go)则使用了sort接口，更加灵活。

#### [链表 (golang)](/go/base/link_list.go)

golang实现的单链表和双链表结构

#### [树 (golang)](/go/base/tree.go)

常用树型结构的golang实现
- 字典树
- 搜索二叉树

#### [cache(golang)](/go/base/lru.go)

- golang实现LRUcache [LRU cache(golang)](/go/base/lru.go)
- golang实现LFUcache [LFU cache(golang)](/go/base/lfu.go)

## 算法部分

> ***做任何算法的时候，都要先弄清需求！如果是需要构造一个函数，那一定要弄清楚函数的调用方式、各参数的含义，多举几个例子说明。只有弄懂了这个函数应该是怎样的，才有可能写出符合要求的函数***

- [数组相关](#array)

- [排序相关](#sort)

- [递归相关](#recursive)

- [其他](#other)

- [leetcode (golang)] (#leetcode)

### Array [array.js](/js/array.js)

- indexOfArray：数组子串位置查询

- arrayFilter：数组筛选

- arrayUnique: 数组去重

- combineArray: 数组归并排序

- longestSubArray: 数组最长无重复子串查找

- longestSubArrayHash: 利用哈希表去重的数组最长无重复子串查找

### sort [sort.js](/js/sort.js)

- [bubbleSort：冒泡排序](/js/sort.js#L20)

- [selectSort：选择排序](/js/sort.js#L36)

- [straightInsertionSort: 直接插入排序](/js/sort.js#L52)

- [shellSort：希尔排序](/js/sort.js#L72)

- [quickSort：快速排序](/js/sort.js#L101)

- [inPlaceQuickSort：原地快速排序](/js/sort.js#L127)

- [mergeSort：归并排序](/js/sort.js#L159)

- [heapSort：堆排序](/js/sort.js#L192)

### recursive[recursive.js](/md/递归.md)

- [阶梯问题](/js/recursive.js#L10)

- [链式函数](/js/recursive.js#L36)

- [汉诺塔问题](/js/recursive.js#L47)

### other

- [订阅发布者事件模型](/js/event.js)

- [节流函数](/js/others.js#L7)

- [防抖函数](/js/others.js#L21)

### leetcode

LeetCode的golang解题：[leetcode](/go/leetcode)







