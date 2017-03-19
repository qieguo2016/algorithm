/**
 * @authors     : qieguo
 * @date        : 2017/3/6 0006
 * @description :
 */

// 节流throttle，多次触发但只执行一部分，（恒时间间距执行）
function throttle(method, threshold, ctx) {
  let timer = null;
  return function () {
    const args = [].slice.call(arguments);
    if (!timer) {
      timer = setTimeout(function () {
        timer = null;
        method.apply(ctx, args);
      }, threshold);
    }
  }
}

// 防抖debounce, 多次触发但只执行一次，（时间差大于阈值才执行）
function debounce(method, threshold, ctx) {
  let timer = null;
  return function () {
    const args = [].slice.call(arguments);
    timer && clearTimeout(timer);
    timer = setTimeout(function () {
      method.apply(ctx, args);
    }, threshold);
  };
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

