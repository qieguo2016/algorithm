/**
 * @authors     : qieguo
 * @date        : 2017/1/2
 * @version     : 1.0
 * @description : 树相关算法（二叉树、二叉查找数、堆）
 */

'use strict';

function BinaryRearchTree(arr) {

  function init() {
    if (!Array.isArray(arr)) {
      console.error(arr + ' is not a Array');
      return false;
    }

  }


}


/**
 * 堆排序
 * 应用：用尺寸为n的最小堆（最大堆）来筛选大数据集里面最大（小）的n个数
 * 流程：
 * 1）先使用数据集里面的前n条数据来构造最小堆
 * 2）遍历数据集，将每个数据与堆顶元素比较，若小于堆顶则抛弃，否则替换掉堆顶
 * 3）调整替换后的堆，继续2）
 *
 * 用有序数组+2分查找也是类似的，不过效率低一些。
 */
