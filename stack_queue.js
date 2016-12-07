/**
 * @authors     : qieguo
 * @date        : 2016/11/29
 * @version     : 1.0
 * @description : js模拟链表
 */

'use strict';
var stack = [];

function pairTest(str) {
	var open = {
		'<': '>',
		'{': '}',
		'(': ')',
	};
	var close = {
		'>': '<',
		'}': '{',
		')': '(',
	};
	var stack = [];
	var result = [];
	for (var i = 0, len = str.length; i < len; i++) {
		if (open[str[i]]) {
			stack.push({index: i, value: str[i]});
		}
		if (close[str[i]]) {
			if (close[str[i]] === stack[stack.length - 1].value) {
				var temp = stack.pop();
				result.push(str.slice(temp.index + 1, i));
			} else {
				throw new Error('匹配出错！');
			}
		}
	}
	return result;
}

console.log(pairTest('sdf\<asdsdfeesf{sdfefi{esadf{aefw}sdfw}sd}\>'));