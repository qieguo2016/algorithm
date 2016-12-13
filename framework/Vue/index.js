/**
 * @authors     : qieguo
 * @date        : 2016/12/12 0012
 * @version     : 1.0
 * @description :
 * @reference   ：https://github.com/DMQ/mvvm#_2， https://github.com/fwing1987/MyVue
 */

'use strict';

// 整体思路：
// 1） Observer实现对vm的监视
// 2） Compiler实现对模板的编译，将vm更新到视图上
// 3） Watcher连接Observer与Compiler，订阅Observer消息后触发视图更新

// Observer
function observe(data) {
  // 设置开始和递归终止条件
  if (!data || typeof data !== 'object') {
    return;
  }
  // 不能直接使用for循环，避开闭包陷阱
  Object.keys(data).forEach(function (key) {
    defineReactive(data, key, data[key]);
  })
}

function defineReactive(data, key, val) {
  var dep = new Dep();
  observe(val);   // 递归对象属性到基本类型为止
  Object.defineProperty(data, key, {
    enumerable: true,    // 枚举
    configurable: false, // 不可再配置
    get: function () {
      console.log('getter');
      // 由于需要在闭包内添加watcher，所以通过Dep定义一个全局target属性，暂存watcher, 添加完移除
      Dep.target && dep.addSub(Dep.target);
      return val;
    },
    set: function (newVal) {
      if (val === newVal) {
        return;
      }
      console.log('setter');
      val = newVal;  // setter本身已经做了赋值，val作为一个闭包变量，保存最新值
      dep.notify();  // 触发通知
    },
  })
}

// dependence
var Dep = function () {
  this.subs = {};
};

Dep.prototype.addSub = function (target) {
  if (!this.subs[target.uid]) {  //防止重复添加
    this.subs[target.uid] = target;
  }
};

Dep.prototype.notify = function (newVal) {
  for (var uid in this.subs) {
    this.subs[uid].update(newVal);
  }
};

// Compiler
function Compile() {

}