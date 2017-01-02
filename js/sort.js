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
function bubbleSort(target) {
  var temp;
  for (var j = target.length; j > 0; j--) {
    for (var i = 0; i < j - 1; i++) {
      if (target[i] > target[i + 1]) {
        temp = target[i];
        target[i] = target[i + 1];
        target[i + 1] = temp;
      }
    }
  }
  return target;
}
function bubbleSortTest() {
  var target = [5, 3, 14, 65, 35, 90, 23];
  console.log('bubbleSort test', bubbleSort(target));
}
// bubbleSortTest()

/**
 * 选择排序
 * @param       : <Array> target数组
 * @description : 一次内循环得到最大值，然后只交换一次次序，将最大值和内循环末尾对调。
 */
function selectSort(target) {
  for (var j = target.length; j > 0; j--) {
    var maxIndex = 0;
    for (var i = 1; i < j; i++) {
      maxIndex = target[maxIndex] > target[i] ? maxIndex : i;
    }
    var temp = target[j - 1];
    target[j - 1] = target[maxIndex];
    target[maxIndex] = temp;
  }
  return target;
}
function selectSortTest() {
  var target = [5, 3, 14, 65, 35, 90, 23];
  console.log('selectSort test', selectSort(target));
}
// selectSortTest()

/**
 * 快速排序
 * @param       : <Array> target数组
 * @description : 选择一个元素将数组分隔成两部分，比该元素小的放该元素前面，比该元素大放后面；
 *                然后递归快速排序，最终得到一个排序后数组
 */
function quickSort(target) {

  if (target.length < 2) { return target; }  // 先定义递归终止条件

  var baseIndex = Math.floor(target.length / 2);
  var left = [];
  var right = [];

  for (var i = 0; i < target.length; i++) {
    if (i === baseIndex) {
      continue;
    }
    if (target[i] < target[baseIndex]) {
      left.push(target[i]);
    } else {
      right.push(target[i]);
    }
  }
  left = quickSort(left)
  right = quickSort(right)
  return left.concat(target[baseIndex], right); // 递归出口
}
function quickSortTest() {
  var target = [5, 3, 14, 65, 35, 90, 23];
  console.log('quickSort test', quickSort(target));
}
// quickSortTest()

/**
 * 归并排序
 * @param       : <Array> target 要归并排序的数组
 * @description : 归并排序，将数组递归分割成两个子数组直至数组只有一个元素，然后将这两个有序数组合并成一个有序数组;
 */
function mergeSortedArray(arrA, arrB) {
  var result = [];
  var i = 0, j = 0, targetLen = arrA.length, toolLen = arrB.length;
  while (i < targetLen && j < toolLen) {
    if (arrA[i] < arrB[j]) {
      result.push(arrA[i++]);
    } else {
      result.push(arrB[j++]);
    }
  }
  while (i < targetLen) {
    result.push(arrA[i++])
  }
  while (j < toolLen) {
    result.push(arrB[j++])
  }
  return result;
}
function mergeSort(target) {
  if (target.length === 1) {
    return target;
  }
  var mid = Math.floor(target.length / 2);
  var left = target.slice(0, mid);
  var right = target.slice(mid);
  return mergeSortedArray(mergeSort(left), mergeSort(right));
}
function mergeSortTest() {
  var target = [1, 25, 100, 8, 90, 11, 10];
  console.log('\nmergeSort test:\n', mergeSort(target));
}
// mergeSortTest()