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

/**
 * 异步操作池
 * @param {Array<Promise>} tasks
 * @param {Number} limit 最大并发数，默认为1表示串行
 * 两种错误处理：1、其中一个任务出错就停止所有；2、执行完所有任务后收集所有错误
 */
function asyncPoolByThunk(tasks, limit = 1, cb = () => {}, options = {}) {
	const stopImmediate = options.stopImmediate;
	let index = limit - 1;
	let errs = [];

	function next(err) {
		if (err) {
			errs.push(err);
			if (stopImmediate) {
				cb(errs);
				return;
			}
		}
		if (index > tasks.length) {
			cb(errs);
			return;
		}
		const current = tasks[++index];
		if (typeof current === 'function') {
			if (stopImmediate && errs.length > 0) {
				return;
			}
			current(next);
		}
	}

	for (let k = 0; k < limit; k++) {
		tasks[k](next);
	}

}

function asyncFn(params) {
	return function (next) {
		setTimeout(() => {
			console.log('async', params);
			next(params  === '4' ? '5555' : null);
		}, 1000);
	}
}
asyncPoolByThunk('123456789'.split('').map(el => asyncFn(el)), 3, function (errs) {
	console.log('errs', errs);
}, { stopImmediate: true });
