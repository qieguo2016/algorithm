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

var methods = []
function countSteps(steps, max, method) {
  method = method || '';
  if(steps === 0) {
    methods.push(method);
    return method;
  }
  if(steps < max) {
    countSteps(steps, steps, method)
  } else {
    for(var i = 1; i <= max; i++) {
      countSteps(steps - i, max, method + String(i));
    }
  }
}

countSteps(4, 3);
console.log('methods', methods);

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

/**
 * 汉诺塔问题
 * 编写函数输出汉诺塔移动轨迹，move(n, a, b, c)，n为盘子数量，a为源头，b为中转柱子，c为目标柱子
 */
function move(n, a, b, c) {
  if(n === 1) {
    console.log(`${a} --> ${c}`);
  } else {
    move(n - 1, a, c, b);
    move(1, a, b, c);
    move(n - 1, b, a, c)
  }
}

// move(5, 'A', 'B', 'C');
