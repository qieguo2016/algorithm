/**
 * @authors     : qieguo
 * @date        : 2016/12/8 0008
 * @version     : 1.0
 * @description : 常见排序算法
 */

'use strict';

/**
 * 交换数组中两元素位置
 * @param       : i, j: 待交换的两元素下标
 */
Array.prototype.swap = function(i, j) {
  const temp = this[i];
  this[i] = this[j];
  this[j] = temp;
}

/**
 * 冒泡排序
 * @param       : <Array> target数组
 * @description : 冒泡排序，更贴切的形容应该是沉底排序，每一轮内循环就是最大数沉底了。
 */
module.exports.bubbleSort = function bubbleSort(target) {
  for(var j = target.length; j > 0; j--) {
    for(var i = 0; i < j - 1; i++) {
      if(target[i] > target[i + 1]) {
        target.swap(i, i + 1);
      }
    }
  }
  return target;
};

/**
 * 选择排序
 * @param       : <Array> target数组
 * @description : 一次内循环得到最大值的下标，然后只交换一次次序，将最大值和内循环末尾对调。
 */
module.exports.selectSort = function selectSort(target) {
  for(var j = target.length; j > 0; j--) {
    var maxIndex = 0;
    for(var i = 1; i < j; i++) {
      maxIndex = target[maxIndex] > target[i] ? maxIndex : i;
    }
    target.swap(maxIndex, j - 1);
  }
  return target;
};

/**
 * 直接插入排序
 * @param       : <Array> target数组
 * @description : 将当前元素与前面元素逐一比较，比前方元素小时则将前方元素后移，直到比前方元素大则落位
 */
module.exports.straightInsertionSort = function straightInsertionSort(target) {
  for(let i = 1; i < target.length; i++) {
    var j = i;
    var base = target[i];
    while(j > 0 && base < target[j - 1]) {
      target[j] = target[j - 1];
      j--;
    }
    if(j < i) {
      target[j] = base;
    }
  }
  return target;
};

/**
 * 希尔排序
 * @param       : <Array> target数组
 * @description : 插入排序的改版：使用定好的偏移量分组，组内进行插入排序；减小偏移量重复进行分组排序直到偏移量为1.
 */
module.exports.shellSort = function shellSort(target) {
  const len = target.length;
  // 偏移量递减
  for(let dx = Math.floor(len / 2); dx > 0; dx = Math.floor(dx / 2)) {
    // 按偏移量分组[0,dx,2dx...], [1, 1+dx, 1+2dx...]
    for(let i = 0; i < dx; i++) {
      // 组内插入排序
      for(let j = i + dx; j < len; j += dx) {
        let k = j;
        let base = target[k];
        // 插入元素
        while(k > i && base < target[k - dx]) {
          target[k] = target[k - dx];
          k -= dx;
        }
        if(k < j) {
          target[k] = base;
        }
      }
    }
  }
  return target;
};

/**
 * 快速排序
 * @param       : <Array> target数组
 * @description : 选择一个元素将数组分隔成两部分，比该元素小的放该元素前面，比该元素大放后面；
 *                然后递归快速排序，最终得到一个排序后数组
 */
module.exports.quickSort = function quickSort(target) {
  // 先定义递归终止条件
  if(target.length < 2) { return target; }

  var baseIndex = 0;
  var left = [];
  var right = [];

  for(var i = 1; i < target.length; i++) {
    if(target[i] < target[baseIndex]) {
      left.push(target[i]);
    } else {
      right.push(target[i]);
    }
  }
  left = quickSort(left)
  right = quickSort(right)
  return left.concat(target[baseIndex], right); // 递归出口
};

/**
 * 原地快速排序
 * @param       : <Array> target
 * @description : 上面的快排每次都开辟一个数组，浪费空间。常规做法是两边查找到中间，两两交换位置
 */
function _inPlaceQuickSort(target, left, right) {
  // 先定义递归终止条件
  if(left >= right) { return target; }

  var base = target[left];
  var i = left;
  var j = right;
  while(i < j) {
    while(i < j && target[j] >= base) {
      j--;
    }
    target[i] = target[j];
    while(i < j && target[i] <= base) {
      i++;
    }
    target[j] = target[i];
  }
  target[i] = base;
  // 函数副作用已经改变了传入数组，但是用显式返回看起来更清晰
  target = _inPlaceQuickSort(target, left, i - 1);
  target = _inPlaceQuickSort(target, i + 1, right);
  return target;
}
module.exports.inPlaceQuickSort = function inPlaceQuickSort(target) {
  return _inPlaceQuickSort(target, 0, target.length - 1)
};

/**
 * 归并排序
 * @param       : <Array> target 要归并排序的数组
 * @description : 归并排序，将数组递归分割成两个子数组直至数组只有一个元素，然后将这两个有序数组合并成一个有序数组;
 */
function mergeSortedArray(arrA, arrB) {
  var result = [];
  var i = 0, j = 0, targetLen = arrA.length, toolLen = arrB.length;
  while(i < targetLen && j < toolLen) {
    if(arrA[i] < arrB[j]) {
      result.push(arrA[i++]);
    } else {
      result.push(arrB[j++]);
    }
  }
  while(i < targetLen) {
    result.push(arrA[i++])
  }
  while(j < toolLen) {
    result.push(arrB[j++])
  }
  return result;
}
module.exports.mergeSort = function mergeSort(target) {
  if(target.length === 1) {
    return target;
  }
  var mid = Math.floor(target.length / 2);
  var left = target.slice(0, mid);
  var right = target.slice(mid);
  return mergeSortedArray(mergeSort(left), mergeSort(right));
};

/**
 * 堆排序
 * @param       : <Array> target
 * @description : 通过构建大(小)根堆的方式进行排序，PS：使用函数副作用来进行原地排序
 */
// 递归调整 i~j 层的大根堆
function adjustMaxHeap(target, i, j) {
  let parent = i;
  let left = 2 * i + 1;
  let right = 2 * i + 2;

  // 比较父节点与左右叶子节点，取最大值的下标设为父节点下标
  if(left < j && target[parent] < target[left]) {
    parent = left;
  }
  if(right < j && target[parent] < target[right]) {
    parent = right;
  }
  // 只有父节点发生改变才会破坏大根堆结构，此时才需要继续调整下级大根堆
  if(parent != i) {
    target.swap(i, parent);
    adjustMaxHeap(target, parent, j);
  }
}
// 构建大根堆就是不断调整最大堆的过程，只要从最后一个父节点往上调整到第一个父节点，就能构建出大根堆
// 从0开始的n层堆的结构：len = 2^n - 1，第n层全是叶子，所以第n-1层的最后一个父节点就是floor(len/2)-1
function buildMaxHeap(target) {
  const len = target.length;
  for(let i = Math.floor(len / 2) - 1; i >= 0; i--) {
    adjustMaxHeap(target, i, len);
  }
}

function sortMaxHeap(target) {
  for(let i = target.length - 1; i > 0; i--) {
    target.swap(0, i);
    adjustMaxHeap(target, 0, i);
  }
}
// 先构建一个大根堆，然后从最后一个元素开始交换堆顶元素，每次交换都调整根堆，直到数组头则完成排序
module.exports.heapSort = function heapSort(target) {
  buildMaxHeap(target);
  sortMaxHeap(target);
  return target;
};

/**
 * 堆排序提取部分记录
 * 从大数据中提取最大(小)的n条记录，也可以用小(大)根堆来实现
 * 先用数据集中前n条数据构造一个小根堆，然后将后面的数据依次通过这个小根堆：
 * 比堆顶小的数据直接丢弃，比堆顶大则替换堆顶，然后调整根堆。最后输出小根堆的排序
 */
module.exports.topSortViaHeap = function topSortViaHeap(target) {
  const len = 10;
  let ret = target.slice(0, len);
  buildMaxHeap(ret);
  for(var i = len; i < target.length; i++) {
    if(target[i] < ret[0]) {
      ret[0] = target[i];
      adjustMaxHeap(ret, 0, ret.length);
    }
  }
  sortMaxHeap(ret);
  return ret;
}
