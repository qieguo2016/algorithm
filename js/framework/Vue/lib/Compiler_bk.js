/**
 * @authors     : qieguo
 * @date        : 2016/12/13
 * @version     : 1.0
 * @description : Compiler，实现对模板的编译，提取指令并将vm与视图关联起来
 */

function Compiler(options) {
	// create node
	this.$el = options.el;

	// save viewModel
	this.vm = options.vm;

	// to documentFragment
	if (this.$el) {
		this.$fragment = nodeToFragment(this.$el);
		this.compile(this.$fragment);
		this.$el.appendChild(this.$fragment);
	}
}

Compiler.prototype = {
	// 编译主体
	compile    : function (node, scope) {
		var self = this;
		if (node.childNodes && node.childNodes.length) {
			[].slice.call(node.childNodes).forEach(function (child) {
				if (child.nodeType === 3) {
					self.compileText(child, scope);
				} else if (child.nodeType === 1) {
					self.compileNode(child, scope);
				}
			});
		}
	},
	// 编译文本元素
	compileText: function (node, scope) {
		scope = scope || this.vm;
		var text = node.textContent.trim();
		if (!text) {
			return;
		}
		var exp = parseTextExp(text);
		// 将exp加入监听列表 node, exp, vm, type
		compileUtil.bindWatcher(node, scope, exp, 'text')
	},
	// 编译节点元素
	compileNode: function (node, scope) {
		var attrs = node.attributes,
			self = this;
		scope = scope || this.vm;
		[].forEach.call(attrs, function (attr) {
			var attrName = attr.name,
				exp = attr.value;
			if (self.isDirective(attrName)) {
				var dir = attrName.substring(2);  // v-dir

				if (self.isEventDirective(dir)) {
					// 事件指令 v-on:click="handle"
					compileUtil.eventHandler(node, scope, exp, dir);

				} else if (self.isAttrDirective(dir)) {
					// 属性指令 v-bind:id="id", v-bind:class="class"
					compileUtil.attrBind(node, scope, exp, dir);

				} else {

					// 普通指令 v-text="variable",v-model="variable"
					var key = dir + 'Handler';
					compileUtil[key] && compileUtil[key](node, scope, exp, dir);

				}
				node.removeAttribute(attrName);
			}
		});
		// @todo 增加编译控制（延迟编译），比如if和for
		self.compile(node, scope);
	},
	// 判断是否指令： v-开头
	isDirective: function (exp) {
		return exp.indexOf('v-') === 0;
	},

	// 事件类型指令，v-on:click="handler"
	isEventDirective: function (exp) {
		return exp.indexOf('on') === 0;
	},

	// HTML属性绑定 v-bind:class="cls"
	isAttrDirective: function (exp) {
		return exp.indexOf('bind') === 0;
	},

	isForDirective: function (exp) {
		return 'for' === exp;
	},
};

// 复制节点到文档碎片
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

// 忽略注释节点和换行节点
function isIgnorable(node) {
	// ignore comment node || a text node
	var regIgnorable = /[^\t\n\r]*/;
	return (node.nodeType == 8) || ((node.nodeType == 3) && (regIgnorable.test(node.textContent)));
}

// 解析文本表达式
function parseTextExp(text) {
	var regText = /\{\{(.+?)\}\}/g;
	var pieces = text.split(regText);
	var matches = text.match(regText);
	// 文本节点转化为常量和变量的组合表达式，PS：表达式中的空格不管，其他空格要保留
	// 'a {{b+"text"}} c {{d+Math.PI}}' => '"a " + b + "text" + " c" + d + Math.PI'
	var tokens = [];
	pieces.forEach(function (piece) {
		if (matches && matches.indexOf('{{' + piece + '}}') > -1) {    // 注意排除无{{}}的情况
			tokens.push(piece);
		} else if (piece) {
			tokens.push('"' + piece + '"');
		}
	});
	return tokens.join('+');
}

// 解析class表达式，@todo 先不带过滤器
function parseClassExp(text) {

}

// 解析style表达式
function parseStyleExp(text) {

}

/**
 * 指令处理，指令主要有：
 * v-text： 表达式编译 @done
 * v-model：数据视图双向绑定 @done
 * v-on：事件绑定 @todo 表达式执行
 * v-bind：控制属性，@todo style和class是特例，区别对待
 * v-show：控制可视化属性，可归在v-bind内
 * v-if、v-for、v-else（暂不做）：控制流，根据当前值会对子元素造成影响：
 * v-html： html编译，要做一定的xss拦截
 * v-pre、v-cloak、v-once：控制不编译、保持内容不变，单次编译暂时不做：
 * */
var compileUtil = {
	// 绑定监听者
	bindWatcher: function (node, scope, exp, dir) {
		//添加一个Watcher，监听exp相关的所有字段变化
		var updateFn = updater[dir] || updater.attr;
		var watcher = new Watcher(exp, scope, function (newVal) {
			updateFn && updateFn(node, newVal, dir);
		});
	},

	// 绑定事件，v-on:click=handle
	eventHandler: function (node, scope, exp, dir) {
		var eventType = dir.split(':')[1], fn = scope[exp];
		if (eventType && fn) {
			node.addEventListener(eventType, fn.bind(scope));  // bind生成一个绑定this的新函数，而call和apply只是调用
		}
	},

	// model双向绑定，v-model
	modelHandler: function (node, scope, exp) {
		if (node.tagName.toLowerCase() === 'input') {
			this.bindWatcher(node, scope, exp, 'value');
			// @FIXME input是高频事件，要做节流?
			node.addEventListener('input', function (e) {
				node.isInputting = true;   // 由于上面绑定了自动更新，循环依赖了，中文输入法不能用。这里加入一个标志避开自动update
				var newValue = e.target.value;
				scope[exp] = newValue;
			});
		}
	},

	htmlHandler: function (node, scope, exp, dir) {
		this.bindWatcher(node, scope, exp, dir);
	},

	textHandler: function (node, scope, exp, dir) {
		this.bindWatcher(node, scope, exp, dir);
	},

	showHandler: function (node, scope, exp, dir) {

	},

	ifHandler: function (node, scope, exp, dir) {

	},

	// 绑定attr，v-bind:class="cls"
	attrBind: function (node, scope, exp, dir) {
		var attr = dir.split(':')[1];
		// @todo 分class、style、普通属性三大类处理，分别求attr
		switch (attr) {
			case 'class':
				exp = parseClassExp(exp);
				break;
			case 'style':
				exp = parseStyleExp(exp);
				break
			default:

		}
		this.bindWatcher(node, scope, exp, attr)
	},

};

var updater = {
	text : function (node, newVal) {
		node.textContent = typeof newVal === 'undefined' ? '' : newVal;
	},
	html : function (node, newVal) {
		node.innerHTML = typeof newVal == 'undefined' ? '' : newVal;
	},
	value: function (node, newVal) {
		// 当有输入的时候循环依赖了，中文输入法不能用。这里加入一个标志避开自动update
		if (!node.isInputting) {
			node.value = newVal ? newVal : '';
		}
		node.isInputting = false;  // 记得要重置标志
	},
	attr : function (node, newVal, attrName) {
		newVal = typeof newVal === 'undefined' ? '' : newVal;
		node.setAttribute(attrName, newVal);
	},
};