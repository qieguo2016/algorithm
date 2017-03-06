/**
 * @authors     : qieguo
 * @date        : 2017/1/2 0002
 * @version     : 1.0
 * @description : 一些递归算法
 */

'use strict';

/**
 * 阶梯问题
 * 假设一个楼梯有 N 阶台阶，人每次最多可以跨 M 阶，求总共的爬楼梯方案数。
 * 例如楼梯总共有3个台阶，人每次最多跨2个台阶，也就是说人每次可以走1个，也可以走2个，但最多不会超过2个，那么楼梯总共有这么几种走法：
 * 111，12，21
 */
function getStepMethodNum(steps, max) {
	var sumStep = 0;
	//总台阶数为0时，终止递归循环
	if (steps === 0) {
		return 1;
	}
	if (steps >= max) {
		//如果steps大于每步最大台阶数，则设置第一步为m之内的一个台阶数，然后递归循环
		for (var i = 1; i <= max; i++) {
			sumStep += getStepMethodNum(steps - i, max);
		}
	} else {
		//如果steps小于m，则将一步最大台阶数缩小为steps，重新递归
		sumStep = getStepMethodNum(steps, steps);
	}
	return sumStep;
}

function getStepMethodNumTest() {
	console.log('\n getStepMethodNumTest: \n', getStepMethodNum(3, 4));
}
getStepMethodNumTest()


/**
 * 链式函数
 * 编写阶乘函数 fn，使得 fn(2)(3) = 6，fn(2)(3)(4) = 24
 * 这里用到递归、函数柯里化、valueOf几个点，注意()运算符是从左到右执行的
*/
function mul(x) {
	const fn = y => mul(x * y);  // 返回一个函数，函数参数里面做乘法运算
	fn.valueOf = () => x;  // 改写valueOf，在链式运算最后一步输出结果
	return fn;
}

