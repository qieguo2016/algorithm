/**
 * @authors     : qieguo
 * @date        : 2017/3/6 0006
 * @description :
 */

// 节流throttle，多次触发但只执行一部分，（恒时间间距执行）
function throttle(fn, delay) {
	var timer = null;
	return function () {
		var ctx = this;
		var args = this.arguments;
		if (!timer) {
			timer = setTimeout(function () {
				timer = null;
				fn.apply(ctx, args);
			}, delay);
		}
	}
}

// 防抖debounce, 多次触发但只执行一次，（时间差大于阈值才执行）
function debounce(fn, delay) {
	var timer = null;
	return function () {
		var ctx = this;
		var args = this.arguments;
		clearTimeout(timer)
		timer = setTimeout(function () {
			fn.apply(ctx, args);
		}, delay);
	}
}

// 轮循函数
// usage: wait(fn.bind(ctx, ...args), 10000);
function wait(fn, timeout, tick) {
	timeout = timeout || 5000;
	tick = tick || 250;
	var timeoutTimer = null;
	var execTimer = null;

	return new Promise(function (resolve, reject) {

		timeoutTimer = setTimeout(function () {
			clearTimeout(execTimer);
			reject(new Error('polling fail because timeout'));
		}, timeout);

		tickHandler(fn);

		function tickHandler(fn) {
			var ret = fn();
			if (!ret) {
				execTimer = setTimeout(function () {
					tickHandler(fn);
				}, tick)
			} else {
				clearTimeout(timeoutTimer);
				resolve();
			}
		}
	});
}

var n = 1;
wait(function () {
	console.log(n++);
	return n > 10;
}, 2000, 300).then(function () {
	console.log('===== end ====')
}).catch(function (err) {
	console.error('error', err);
});

