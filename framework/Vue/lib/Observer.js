/**
 * @authors     : qieguo
 * @date        : 2016/12/13
 * @version     : 1.0
 * @description : Observer
 */

'use strict';

function Observer(data) {
	this.data = data;
	this.observe(data);
}

// Observer
Observer.prototype.observe = function (data) {
	var self = this;
	// 设置开始和递归终止条件
	if (!data || typeof data !== 'object') {
		return;
	}
	// 不能直接使用for循环，避开闭包陷阱
	Object.keys(data).forEach(function (key) {
		self.defineReactive(data, key, data[key]);
	})
}

Observer.prototype.defineReactive = function (data, key, val) {
	var dep = new Dep();
	var self = this;
	self.observe(val);   // 递归对象属性到基本类型为止
	Object.defineProperty(data, key, {
		enumerable  : true,    // 枚举
		configurable: false, // 不可再配置
		get         : function () {
			console.log('getter');
			// 由于需要在闭包内添加watcher，所以通过Dep定义一个全局target属性，暂存watcher, 添加完移除
			Dep.target && dep.addSub(Dep.target);
			return val;
		},
		set         : function (newVal) {
			if (val === newVal) {
				return;
			}
			console.log('setter');
			val = newVal;  // setter本身已经做了赋值，val作为一个闭包变量，保存最新值
			self.observe(newVal);
			dep.notify(newVal);  // 触发通知
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