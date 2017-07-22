/**
 * @authors     : qieguo
 * @date        : 2016/12/8 0008
 * @version     : 1.0
 * @description : 常见排序算法
 */

'use strict';

/**
 * 冒泡排序
 * @param       : <Array> target数组
 * @description : 冒泡排序，更贴切的形容应该是沉底排序，每一轮内循环就是最大数沉底了。
 */
module.exports.bubbleSort = function bubbleSort(target) {
  var temp;
  for(var j = target.length; j > 0; j--) {
    for(var i = 0; i < j - 1; i++) {
      if(target[i] > target[i + 1]) {
        temp = target[i];
        target[i] = target[i + 1];
        target[i + 1] = temp;
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
    var temp = target[j - 1];
    target[j - 1] = target[maxIndex];
    target[maxIndex] = temp;
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
module.exports.inPlaceQuickSort = function inPlaceQuickSort(target) {
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
    // 也可以利用函数副作用改变传入数组，但是用显式返回更清晰
    target = _inPlaceQuickSort(target, left, i - 1);
    target = _inPlaceQuickSort(target, i + 1, right);
    return target;
  }

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


