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
 * @description : 数组、字符串的子集位置查询
 */

function indexOfArray(target, tool) {
  for (var i = 0, targetLen = target.length; i < targetLen; i++) {
    var isEqual = true;
    for (var j = 0, toolLen = tool.length; j < toolLen; j++) {
      if (target[i + j] !== tool[j]) {
        isEqual = false;
        break;
      }
    }
    if (isEqual) {
      return i;
    }
  }
  return -1;
}
// console.log(indexOfArray('abcdedfg', 'df'));


/**
 * 根据筛选条件从目标数组中返回符合条件的子集
 * @param       : <Array> target 要筛选的数据（数组）
 * @param       : <Object> tool  筛选条件（对象）
 * @return      : 返回符合条件的子集数组，否则返回空数组
 * @description : 数组、字符串的子集位置查询
 */

function arrayFilter(target, tool) {
  return target.filter(function (item) {
    var keep = true;
    for (var key in tool) {
      if (typeof item[key] === 'undefined' || item[key] !== tool[key]) {
        keep = false;
        break;
      }
    }
    return keep;
  });
}
// var target = [
//   {name: 'Jack', age: 18, sex: 'f'},
//   {name: 'Jack', age: 20},
//   {name: 'Mike', age: 25}
// ];
// // var tool = {num: 20}
// // var tool = {}
// console.log(dataFilter(target, tool));


/**
 * 数组去重
 * @param       : <Array> target 要去重的数组
 * @return      : 返回去重后的子集数组
 * @description : 数组去重
 */
function arrayUnique(target) {
  var result = [target[0]];
  for (var i = 1, targetLen = target.length; i < targetLen; i++) {
    var isUnique = true;
    for (var j = 0, resultLen = result.length; j < resultLen; j++) {
      if (result[j] === target[i]) {
        isUnique = false;
        break
      }
    }
    if (isUnique) {
      result.push(target[i]);
    }
  }
  return result;
}
var target = [1, 2, 3, 3, 2, '3', {},{}];
console.log(arrayUnique(target));