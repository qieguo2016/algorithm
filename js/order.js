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
function bubble(target) {
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
console.log('bubble test', bubble([5, 3, 14, 65, 35, 90, 23]));

/**
 * 选择排序
 * @param       : <Array> target数组
 * @description : 选择排序，一次内循环只交换一次次序。
 */
function select(target) {
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
console.log('select test', select([5, 3, 14, 65, 35, 90, 23]));
