/**
 * @authors     : qieguo
 * @date        : 2016/12/13
 * @version     : 1.0
 * @description : Compiler
 */

'use strict';

var regIgnorable = /[^\t\n\r]*/;

function Compiler(options) {
	// create node
	this.$el = typeof options.el === 'string'
		? document.querySelector(options.el)
		: options.el || document.createElement('div');

	// to documentFragment
	if (this.$el) {
		this.$fragment = nodeToFragment(this.$el);
		this.compile(this.$fragment);
		this.$el.appendChild(this.$fragment);
	}
}

function nodeToFragment(node) {
	var fragment = document.createDocumentFragment(), child;
	while (child = node.firstChild) {
		if (isIgnorable(child)) {     // delete '\n'
			node.removeChild(child);
		} else {
			fragment.appendChild(child);
		}
	}
	return fragment;
}

function isIgnorable(node) {
	// A comment node || a text node
	return (node.nodeType == 8) || ((node.nodeType == 3) && (regIgnorable.test(node.textContent)));
}

Compiler.prototype.compile = function (node) {
	var self = this;
	if (node.childNodes && node.childNodes.length) {
		[].slice.call(node.childNodes).forEach(function (child) {
			if (child.nodeType === 3) {
				self.compileText(child);
			} else if (child.nodeType === 1) {
				self.compileNode(child);
				self.compile(child);
			}
		});
	}
}

Compiler.prototype.compileText = function (node) {
	console.log('compileText', node.textContent);

	var text = node.textContent.trim();
	var tokens = [];

	var regText = /\{\{(.+?)\}\}/g;
	var pieces = text.split(regText);
	var matches = text.match(regText);

	// 文本节点转化为常量和变量的组合表达式
	// 'a {{b+"text"}} c {{d+Math.PI}}' => '"a " + b + "text" + " c" + d + Math.PI'
	pieces.forEach(function (piece) {
		if (matches.indexOf('{{' + piece + '}}') > -1) {
			tokens.push('(' + piece + ')');
		} else if (piece) {
			tokens.push('"' + piece + '"');
		}
	});

	var exp = tokens.join('+');

	// 将exp加入监控
	console.log('exp', exp);


}

Compiler.prototype.compileNode = function (node) {
	var attrs = node.attributes,
		self = this;
	[].forEach.call(attrs, function (attr) {
		var attrName = attr.name,
			exp = attr.value;
		if (self.isDirective(attrName)) {
			var dir = attrName.substring(2);
			if (self.isEventDirective(dir)) {
				// 事件指令 v-on:click="handle"
				compileUtil.eventHandler(node, self.$vm, exp, dir);
			} else {
				// 普通指令 v-text="variable",v-model,v-class. @todo (v-bind:id?)
				compileUtil[dir] && compileUtil[dir](node, self.$vm, exp);
			}
			node.removeAttribute(attrName);
		}
	});
}

// 判断是否指令： v-开头
Compiler.prototype.isDirective = function (attr) {
	return attr.indexOf('v-') === 0;
}

// 事件类型指令，v-on
Compiler.prototype.isEventDirective = function (attr) {
	return attr.indexOf('on') === 0;
}

// 指令处理
var compileUtil = {
	text: function (node, vm, exp) {
		this.bind(node, vm, exp, 'text');
	},

	html: function (node, vm, exp) {
		this.bind(node, vm, exp, 'html');
	},

	model: function (node, vm, exp) {
		this.bind(node, vm, exp, 'model');

		var me = this,
			val = this._getVMVal(vm, exp);
		node.addEventListener('input', function (e) {
			var newValue = e.target.value;
			if (val === newValue) {
				return;
			}

			me._setVMVal(vm, exp, newValue);
			val = newValue;
		});
	},

	class: function (node, vm, exp) {
		this.bind(node, vm, exp, 'class');
	},

	bind: function (node, vm, exp, dir) {
		var updaterFn = updater[dir + 'Updater'];

		updaterFn && updaterFn(node, this._getVMVal(vm, exp));

		new Watcher(vm, exp, function (value, oldValue) {
			updaterFn && updaterFn(node, value, oldValue);
		});
	},

	// 绑定事件，v-on:click=handle
	eventHandler: function (node, vm, exp, dir) {
		var eventType = dir.split(':')[1],
			fn = vm.$options.methods && vm.$options.methods[exp];
		if (eventType && fn) {
			node.addEventListener(eventType, fn.bind(vm));  // bind生成一个绑定this的新函数，而call和apply只是调用
		}
	},

	_getVMVal: function (vm, exp) {
		var val = vm._data;
		exp = exp.split('.');
		exp.forEach(function (k) {
			val = val[k];
		});
		return val;
	},

	_setVMVal: function (vm, exp, value) {
		var val = vm._data;
		exp = exp.split('.');
		exp.forEach(function (k, i) {
			// 非最后一个key，更新val的值
			if (i < exp.length - 1) {
				val = val[k];
			} else {
				val[k] = value;
			}
		});
	}
};

var updater = {
	textUpdater: function (node, value) {
		node.textContent = typeof value == 'undefined' ? '' : value;
	},

	htmlUpdater: function (node, value) {
		node.innerHTML = typeof value == 'undefined' ? '' : value;
	},

	classUpdater: function (node, value, oldValue) {
		var className = node.className;
		className = className.replace(oldValue, '').replace(/\s$/, '');

		var space = className && String(value) ? ' ' : '';

		node.className = className + space + value;
	},

	modelUpdater: function (node, value, oldValue) {
		node.value = typeof value == 'undefined' ? '' : value;
	}
};