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
// var target = [
//   {name: 'Jack', age: 18, sex: 'f'},
//   {name: 'Jack', age: 20},
//   {name: 'Mike', age: 25}
// ];
// var tool = {age: 20}
// // var tool = {}
// console.log(arrayFilter(target, tool));


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
// var target = [1, 2, 3, 3, 2, '3', {}, {}];
// console.log(arrayUnique(target));

/**
 * 数组归并排序
 * @param       : <Array> target 要归并排序的数组
 * @param       : <Array> tool 要归并排序的数组
 * @return      : 返回合并排序后的数组
 * @description : 归并排序，将两个已经排序的数组合并成一个数组;
 */

function combineArray(target, tool) {
	var result = [];
	var i = 0, j = 0, targetLen = target.length, toolLen = tool.length;
	while (i < targetLen && j < toolLen) {
		if (target[i] < tool[j]) {
			result.push(target[i++]);
		} else {
			result.push(tool[j++]);
		}
	}
	while (i < targetLen) {
		result.push(target[i++])
	}
	while (j < toolLen) {
		result.push(tool[j++])
	}
	return result;
}
// var target = [1, 5, 11, 18, 25, 40, 100, 120];
// var tool = [3, 6, 11, 30, 31, 80, 90, 97];
// console.log(combineArray(target, tool));

