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


