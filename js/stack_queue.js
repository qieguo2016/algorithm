/**
 * @authors     : qieguo
 * @date        : 2016/11/29
 * @version     : 1.0
 * @description : js模拟链表
 */

'use strict';

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

function waterfall(tasks, callback) {

	// var err = null;
	// var res = null;
	// var current;
	// while (fnArray.length !== 0) {
	// 	current = fnArray.shift();
	// 	current.apply(null, args);
	// }

	if (!tasks.length) return callback();
	var taskIndex = 0;

	function nextTask(args) {
		if (taskIndex === tasks.length) {
			return callback.apply(null, [null].concat(args));
		}

		var taskCallback = onlyOnce(rest(function(err, args) {
			if (err) {
				return callback.apply(null, [err].concat(args));
			}
			nextTask(args);
		}));

		args.push(taskCallback);

		var task = tasks[taskIndex++];
		task.apply(null, args);
	}

	nextTask([]);

}