/**
 * @authors     : qieguo
 * @date        : 2016/11/24
 * @version     : 1.0
 * @description : 数组、字符串相关；逻辑实现，不使用API等现成接口
 */


/**
 * 数组、字符串的子集位置查询
 * @param       : <Array>||<String> target,目标字符串
 * @param       : <Array>||<String> tool,子串
 * @return      : <Number> 返回子串在目标中的位置，目标中找不到子串则返回-1
 * @description : 数组、字符串的子集位置查询；可转成字符串用正则匹配，
 */
function indexOfArray(target, tool) {
  for (var i = 0, targetLen = target.length; i < targetLen; i++) {
    for (var j = 0, toolLen = tool.length; j < toolLen; j++) {
      if (target[i + j] !== tool[j]) {
        break;    // j++是在循环体结束后自增，使用break则在自增前就跳出循环了
      }
    }
    if (j === toolLen) {
      return i;
    }
  }
  return -1;
}
function indexOfArrayTest() {
  console.log('indexOfArray test: ', indexOfArray('abcdedfg', 'e'));
}


/**
 * 根据筛选条件从目标数组中返回符合条件的子集
 * @param       : <Array> target 要筛选的数据（数组）
 * @param       : <Object> tool  筛选条件（对象）
 * @return      : 返回符合条件的子集数组，否则返回空数组
 * @description : 数组、字符串的子集位置查询；可用filter高阶函数代替循环
 */
function arrayFilter(target, tool) {
  var result = [];
  for (var i = 0, len = target.length; i < len; i++) {
    var keep = true;
    for (var key in tool) {
      if (typeof target[i][key] === 'undefined' || target[i][key] !== tool[key]) {
        keep = false;
        break;
      }
    }
    if (keep) {
      result.push(target[i]);
    }
  }
  return result;
}
function arrayFilterTest() {
  var target = [
    {name: 'Jack', age: 18, sex: 'f'},
    {name: 'Jack', age: 20},
    {name: 'Mike', age: 25}
  ];
  console.log('\narrayFilter test: \n', arrayFilter(target, {age: 20}));
  console.log('\narrayFilter test: \n', arrayFilter(target, {num: 10}));
  console.log('\narrayFilter test: \n', arrayFilter(target, {}));
}


/**
 * 数组去重
 * @param       : <Array> target 要去重的数组
 * @description : 数组去重；可用es6 set、正则、sort等实现
 */
function arrayUnique(target) {
  // 纯数组硬比较方式，这里要注意是否需要空对象{}的去重，暂不处理
  // var result = [target[0]];
  // for (var i = 1, targetLen = target.length; i < targetLen; i++) {
  //   for (var j = 0, resultLen = result.length; j < resultLen; j++) {
  //     if (result[j] === target[i]) {
  //       break;    // j++是在循环体结束后自增，使用break则在自增前就跳出循环了
  //     }
  //   }
  //   if (j === resultLen) {
  //     result.push(target[i]);
  //   }
  // }

  // 对于去重这种无序的集合，可使用js对象的哈希特性来提高效率，但无法直接区分数字、字符，统一转为字符了
  // Note: 数据量少的情况下，哈希算法本身的复杂度就超过了循环对比，所以性能上反而更差
  var result = [target[0]];
  var temp = Object.create(null);
  temp[target[0]] = {};
  temp[target[0]][(typeof target[0])] = 1;
  // 要区分数字、字符、布尔值、null等类型，必须保存 temp[target[i]][(typeof target[i])] 作为标志
  for (var i = 1, targetLen = target.length; i < targetLen; i++) {
    if (typeof temp[target[i]] === 'undefined' || !temp[target[i]].hasOwnProperty(typeof target[i])) {
      result.push(target[i]);
      temp[target[i]] = {};
      temp[target[i]][(typeof target[i])] = 1;
    }
  }
  return result;
}
function arrayUniqueTest() {
  // var target = [1, 2, 3, 3, '3', '3', 'length', '__proto__', 'prototype', true, false, true, {}, {}, null, null];
  var target = [1, 2, 3, 3, '3', '3', '__proto__', '__proto__', '__proto__', 'prototype', 'prototype', true, false, true, {}, {}, null, null];
  // var target = [1, '1', true, 'true'];
  console.log('\narrayUnique test:\n', arrayUnique(target));
}


/**
 * 合并两个有序数组
 * @param       : <Array> arrA 要合并的有序数组
 * @param       : <Array> arrB 要合并的有序数组
 * @description : 将两个已经排序的数组合并成一个数组;
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
function mergeSortedArrayTest() {
  var target = [1, 5, 11, 18, 25, 40, 100, 120];
  var tool = [3, 6, 11, 30, 31, 80, 90, 97];
  console.log('\ncombineArray test:\n', mergeSortedArray(target, tool));
}


/**
 * 数组最长无重复子串查找
 * @param       : <Array> target 要查找的数组
 * @description : 查找没有重复的最长子串，若使用哈希表判断重复的话，就要重新定位，可以将哈希的value=数组元素当前序号
 */
function longestSubArray(target) {
  var lastStart, lastLen, maxStart, maxLen;
  lastStart = maxStart = 0;
  lastLen = maxLen = 1;

  for (var i = 1, len = target.length; i < len; i++) {
    var noRepeat = true;
    for (var j = lastStart; j < lastStart + lastLen; j++) {
      if (target[i] === target[j]) {
        noRepeat = false;
        if (maxLen < lastLen) {
          maxLen = lastLen;
          maxStart = lastStart;
        }
        i = lastStart = i - (lastStart + lastLen - j) + 1;
        lastLen = 1;
        break;
      }
    }
    if (noRepeat) {
      lastLen++;
    }
  }
  if (maxLen < lastLen) {
    maxLen = lastLen;
    maxStart = lastStart;
  }
  return target.slice(maxStart, maxLen + maxStart);
}
// 哈希版本，js实现起来简单、在数据量大的情况下会有优势
function longestSubArrayHash(target) {
  var lastStart, lastLen, maxStart, maxLen, last = {};
  lastStart = maxStart = 0;
  lastLen = maxLen = 1;
  last[target[lastStart]] = lastStart;

  for (var i = 1, len = target.length; i < len; i++) {
    if (typeof last[target[i]] === 'undefined') {
      lastLen++;
      last[target[i]] = i;
    } else {
      if (maxLen < lastLen) {
        maxLen = lastLen;
        maxStart = lastStart;
      }
      i = lastStart = last[target[i]] + 1;
      last = {};
      last[target[lastStart]] = lastStart;
      lastLen = 1;
    }
  }
  if (maxLen < lastLen) {
    maxLen = lastLen;
    maxStart = lastStart;
  }
  return target.slice(maxStart, maxLen + maxStart);
}
function longestSubArrayTest() {
  var target = [1, 2, 3, 4, 3, 6, 8, 9, 10, 14, 15, 8, 9];
  console.log('\nlongestSubArray test:\n', longestSubArray(target));
  console.log('\nlongestSubArrayHash test:\n', longestSubArrayHash(target));
}


/**
 * 数组重复次数最多的子串查找
 * @param       : <Array> target 要查找的数组
 * @description : 查重，输出重复次数最多的元素及其重复次数。
 */
function countSubArray(target) {
  var sub = {};
  var max = {
    num: 1,
    index: 0
  };
  for (var i = 0, len = target.length; i < len; i++) {
    if (!sub.hasOwnProperty(target[i])) {
      sub[target[i]] = {
        num: 1,
        index: i
      };
    } else {
      var current = sub[target[i]];
      current.num++;
      if (max.num < current.num) {
        max.num = current.num;
        max.index = current.index;
      }
    }
  }
  return {
    element: target[max.index],
    index: max.index,
    count: max.num
  };
}
function countSubArrayTest() {
  var target = [2, 2, 2, 4, 4, 11, 11, 5, 15, 11, 17, 11, 80, 11];
  var target1 = ['qwe', 'as', 'dsfw', 'as', 'kou', 'lpi', 'as', 'jei', 'as'];
  console.log('\ncountSubArrayTest test:\n', countSubArray(target1));
};


/**
 * 查询字符串第一个不重复字母
 * @param       : <Array> target 要查找的数组
 * @description : 查询给定字符串“abcba”，处理得到第一个不重复字母，c
 */
function queryFirstUniqueItem(target) {
  var items = {};
  var last = [];
  for (var i = 0, l = target.length; i < l; i++) {
    var item = target[i];
    if (!items.hasOwnProperty(item)) {
      last.push(i);
      items[item] = i;
    } else {
      var index = last.indexOf(items[item]);
      if (index >= 0) {
        last.splice(index, 1);
      }
    }
  }
  var ret = '';
  if (last.length) {
    ret = target[last[0]];
  }
  return ret
}
function queryFirstUniqueItemTest() {
  var a = "abdcbadegceg";
  var ret = queryFirstUniqueItem(a);
  console.log('ret: ', ret);
  console.log('ret: ', !ret);
}

// 给一个数组如：[[“a”,”b”,”c”],[“d”,”e”],…..]得到[ad,ae,bd,be,cd,ce]
function mapConcat(target) {
  var res = target.reduce(function (pre, next) {
    var ret = [];
    pre.forEach(function (preItem) {
      next.forEach(function (nextItem) {
        if (preItem !== nextItem) {  // 去掉aa这种情况
          ret.push(preItem + nextItem);
        }
      })
    })
    return ret;
  })
  return res;
}
function mapConcatTest() {
  // var target = [['a', 'b', 'c'], ['d', 'e'], ['f', 'g', 'h'], ['j', 'm']];
  var target = [['a', 'b', 'c'], ['a', 'b']];
  console.log('\nmapConcatTest:\n', mapConcat(target));
}
mapConcatTest()