/**
 * @authors     : qieguo
 * @date        : 2016/12/15
 * @version     : 1.0
 * @description : MVVM，框架入口，各个模块的容器，data的代理。
 */

function MVVM(options) {
	this.$data = options.data || {};
	this.$el = typeof options.el === 'string'
		? document.querySelector(options.el)
		: options.el || document.body;
	this.$options = options;

	// 代理属性，直接用vm.variable访问data、method、computed内数据/方法
	this._proxy(options);
	this._proxyMethod(options.methods);   // method不劫持getter/setter

	var ob = new Observer(this.$data);

	if (!ob) return;
	new Compiler({el: this.$el, vm: this});
}

MVVM.prototype = {
	// 代理属性，直接用vm.variable访问data、computed内数据/方法
	_proxy      : function (data) {
		var self = this;
		var proxy = ['data', 'computed'];
		proxy.forEach(function (item) {
			Object.keys(data[item]).forEach(function (key) {
				Object.defineProperty(self, key, {
					configurable: false,
					enumerable  : true,
					get         : function () {
						return self.$data[key] || self.$options.computed[key].call(self);
					},
					set         : function (newVal) {
						if (self.$data.hasOwnProperty(key)) {
							self.$data[key] = newVal;
						} else if (self.$options.computed.hasOwnProperty(key)) {
							self.$options.computed[key] = newVal;
						}
					}
				});
			})
		})
	},
	// method不劫持getter/setter，直接引用
	_proxyMethod: function (methods) {
		var self = this;
		Object.keys(methods).forEach(function (key) {
			self[key] = self.$options.methods[key];
			/*	Object.defineProperty(self, key, {
			 configurable: false,
			 enumerable  : true,
			 get         : function () {
			 return self.$options.methods[key];
			 },
			 set         : function (newFn) {
			 if (typeof newFn === 'function') {
			 self.$options.methods[key] = newFn;
			 }
			 }
			 });*/
		})
	}
}