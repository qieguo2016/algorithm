/**
 * @authors     : qieguo
 * @date        : 2016/12/13
 * @version     : 1.0
 * @description : Compiler
 */

function Compiler(options) {
  // create node
  this.$el = typeof options.el === 'string'
    ? document.querySelector(options.el)
    : options.el || document.body;

  // save viewModel
  this.$vm = options.vm;

  // to documentFragment
  if (this.$el) {
    this.$fragment = nodeToFragment(this.$el);
    this.compile(this.$fragment);
    this.$el.appendChild(this.$fragment);
  }
}

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

Compiler.prototype = {
  // 编译主体
  compile: function (node) {
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
  },
  // 编译文本元素
  compileText: function (node) {

    var text = node.textContent.trim();
    var regText = /\{\{(.+?)\}\}/g;
    var pieces = text.split(regText);
    var matches = text.match(regText);

    // 文本节点转化为常量和变量的组合表达式，PS：表达式中的空格不管，其他空格要保留
    // 'a {{b+"text"}} c {{d+Math.PI}}' => '"a " + b + "text" + " c" + d + Math.PI'
    var tokens = [];
    pieces.forEach(function (piece) {
      if (matches.indexOf('{{' + piece + '}}') > -1) {
        tokens.push(piece);
      } else if (piece) {
        tokens.push('"' + piece + '"');
      }
    });

    // 计算exp刷新文本
    var exp = tokens.join('+');
    node.textContent = this.parseExpression(exp);
    // 将exp加入监听列表 node, exp, vm, type
    compileUtil.bindWatcher(node, exp, self.$vm, 'text')
  },
  // 编译节点元素
  compileNode: function (node) {
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
        } else if (self.isAttrDirective(dir)) {
          // compileUtil.attrBind(node, self.$vm, exp, dir);
        } else {
          // 普通指令 v-text="variable",v-model="variable"
          compileUtil[dir] && compileUtil[dir](node, self.$vm, exp);
        }
        node.removeAttribute(attrName);
      }
    });
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
  // 解析表达式
  parseExpression: function (exp) {
    with (this.$vm.$data) {
      return eval(exp);
    }
  },
};


// 指令处理
var compileUtil = {
  // model双向绑定，v-model
  model: function (node, vm, exp) {
    if (node.tagName.toLowerCase() === 'input') {
      node.addEventListener('input', function (e) {
        // input是高频事件，要做节流
        var newValue = e.target.value;
        vm.data[exp] = newValue;
      });
    }

  },
  // 绑定事件，v-on:click=handle
  eventHandler: function (node, vm, exp, dir) {
    var eventType = dir.split(':')[1],
      fn = vm.$options.methods && vm.$options.methods[exp];
    if (eventType && fn) {
      node.addEventListener(eventType, fn.bind(vm));  // bind生成一个绑定this的新函数，而call和apply只是调用
    }
  },

  // 绑定监听者
  bindWatcher: function (node, exp, vm, type) {
    //绑定view与model
    //添加一个Watcher，监听exp相关的所有字段变化，具体方法可以看Watcher的注释
    var updateFn = updater[type];
    var watcher = new Watcher(exp, vm, function (newVal, oldVal) {
      updateFn && updateFn(node, newVal, oldVal);
    });
  },
  //////////////////////////////////////////////

  // 绑定attr，v-bind:class="cls"
  attrBind: function (node, vm, exp, dir) {
    var attr = dir.split(':')[1];
  },

};


var updater = {
  text: function (node, newVal) {
    node.textContent = typeof newVal === 'undefined' ? '' : newVal;
  },
  html: function (node, newVal) {
    node.innerHTML = typeof newVal == 'undefined' ? '' : newVal;
  },
  value: function (node, newVal) {
    node.value = newVal ? newVal : '';
  },
  attr: function (node, newVal, attrName) {
    node[attrName] = typeof newVal == 'undefined' ? '' : newVal;
  }
};