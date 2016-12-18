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
  compile: function (node, scope) {
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
    var forExp = '';
    [].forEach.call(attrs, function (attr) {
      var attrName = attr.name,
        exp = attr.value;

      // getDirective @return dir
      // switch dir >> go to handler

      if (self.isDirective(attrName)) {
        var dir = attrName.substring(2);  // v-dir
        // for最后编译
        if (self.isForDirective(dir)) {
          forExp = exp;
        } else if (self.isEventDirective(dir)) {
          // 事件指令 v-on:click="handle"
          compileUtil.eventHandler(node, scope, exp, dir);
        } else {
          // v-text="variable",v-bind:id="id"
          var key = dir.split(':')[0] + 'Handler';
          compileUtil[key] && compileUtil[key](node, scope, exp, dir);
        }
        node.removeAttribute(attrName);
      }
    });

    // for最后编译
    if (forExp) {
      self.forHandler(node, scope, forExp)
    } else {
      self.compile(node, scope);
    }
  },

  // 绑定监听者
  bindWatcher: function (node, scope, exp, dir, prop) {
    //添加一个Watcher，监听exp相关的所有字段变化
    var updateFn = updater[dir];
    var watcher = new Watcher(exp, scope, function (newVal) {
      updateFn && updateFn(node, newVal, prop);
    });
  },

  // 各种指令编译处理方法

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
      node.addEventListener('input', function (e) {
        node.isInputting = true;   // 由于上面绑定了自动更新，循环依赖了，中文输入法不能用。这里加入一个标志避开自动update
        var newValue = e.target.value;
        scope[exp] = newValue;
      });
    }
  },

  // @FIXME 变更需要重新编译子元素
  htmlHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, dir);
  },

  textHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, dir);
  },

  showHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, 'style', 'display')
  },

  ifHandler: function (node, scope, exp, dir) {

  },

  // 属性指令 v-bind:id="id", v-bind:class="cls"
  bindHandler: function (node, scope, exp, dir) {
    var attr = dir.split(':')[1];
    switch (attr) {
      case 'class':
        // 拼成 "baseCls "+(a?"acls ":"")+(b?"bcls ":"")的形式
        exp = '"' + node.className + ' "+' + parseClassExp(exp);
        break;
      case 'style':
        // style可以使用style.cssText/node.setAttribute('style','your style')全量更新，也可以使用style.prop单个更新
        // 全量更新只需要监听全量表达式即可，但是初次编译之后其他地方脚本改了propB的话，下一次更新propA也会使用vm的值去覆盖更改后的propB
        // 单个更新的话需要监听多个值，但是不同样式之间无影响，比如初次编译后脚本更改了propC，下一次更新propB是不会影响到propC的
        // 这里使用全量更新，样式写法是这样的：<div v-bind:style="{ color: activeColor, font-size: fontSize }"></div>
        var styleStr = node.getAttribute('style');
        exp = '"' + styleStr + ';"+' + parseStyleExp(exp);
        break;
      default:

    }
    this.bindWatcher(node, scope, exp, 'attr', attr)
  },

  forHandler: function (node, scope, exp, options) {
    var self = this;
    var itemName = exp.split('in')[0].replace(/\s/g, '')
    var arrNames = exp.split('in')[1].replace(/\s/g, '').split('.');
    var arr = scope[arrNames[0]];
    if (arrNames.length === 2) {
      arr = arr[arrNames[1]];
    }
    var parentNode = node.parentNode;
    arr.forEach(function (item) {
      var cloneNode = node.cloneNode(true);
      parentNode.insertBefore(cloneNode, node);
      var forScope = Object.create(scope);  // 注意每次循环要生成一个新对象
      forScope[itemName] = item;
      self.compile(cloneNode, forScope);  // @FIXME 同样的编译应该有缓存机制
    });
    parentNode.removeChild(node);   // 去掉原始模板
  },

  // 判断是否指令： v-开头
  isDirective: function (exp) {
    return exp.indexOf('v-') === 0;
  },
  // 事件类型指令，v-on:click="handler"
  isEventDirective: function (exp) {
    return exp.indexOf('on') === 0;
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
      fragment.appendChild(child);   // 移动操作，将child从原位置移动添加到fragment
    }
  }
  return fragment;
}

// 忽略注释节点和换行节点
function isIgnorable(node) {
  // ignore comment node || a text node
  var regIgnorable = /^[\t\n\r]+/;
  return (node.nodeType == 8) || ((node.nodeType == 3) && (regIgnorable.test(node.textContent)));
}

// 解析文本表达式 @todo 未包含pipe语法
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
      tokens.push('`' + piece + '`');
    }
  });
  return tokens.join('+');
}

// 解析class表达式，@todo 目前未写数组语法
// <div class="static" v-bind:class="{ active: isActive, 'text-danger': hasError }"> </div>
function parseClassExp(exp) {
  if (!exp) {
    return;
  }
  var regObj = /\{(.+?)\}/g;
  var regArr = /\[(.+?)\]/g;
  var result = [];
  if (regObj.test(exp)) {
    var subExp = exp.replace(/[\s\{\}]/g, '').split(',');
    subExp.forEach(function (sub) {
      var key = '"' + sub.split(':')[0].replace(/['"`]/g, '') + ' "';
      var value = sub.split(':')[1];
      result.push('(' + value + '?' + key + ':"")')
    });
  } else if (regArr.test(exp)) {
    var subExp = exp.replace(/[\s\[\]]/g, '').split(',');
  }
  return result.join('+');  // 拼成 (a?"acls ":"")+(b?"bcls ":"")的形式
}

// 解析style表达式 @todo 目前未写数组语法
// <div v-bind:style="{ color: activeColor, font-size: fontSize }"></div>
function parseStyleExp(exp) {
  if (!exp) {
    return;
  }
  var regObj = /\{(.+?)\}/g;
  var regArr = /\[(.+?)\]/g;
  var result = [];
  if (regObj.test(exp)) {
    var subExp = exp.replace(/[\s\{\}]/g, '').split(',');
    subExp.forEach(function (sub) {
      // "color:"activeColor;"font-size:"fontSize;
      var key = '"' + sub.split(':')[0].replace(/['"`]/g, '') + ':"+';
      var value = sub.split(':')[1];
      result.push(key + value + '+";"');
    });
  } else if (regArr.test(exp)) {
    var subExp = exp.replace(/[\s\[\]]/g, '').split(',');
  }
  return result.join('+');  // 拼成 (a?"acls ":"")+(b?"bcls ":"")的形式
}

/**
 * 指令处理，指令主要有：
 * v-text： 表达式编译 @done
 * v-model：数据视图双向绑定 @done
 * v-on：事件绑定 @done
 * v-bind：控制属性
 * v-show：控制可视化属性，可归在v-bind内
 * v-if、v-for、v-else（暂不做）：控制流，根据当前值会对子元素造成影响：
 * v-html： html编译，要做一定的xss拦截
 * v-pre、v-cloak、v-once：控制不编译、保持内容不变，单次编译暂时不做：
 * */
var compileUtil = {
  // 绑定监听者
  bindWatcher: function (node, scope, exp, dir, prop) {
    //添加一个Watcher，监听exp相关的所有字段变化
    var updateFn = updater[dir];
    var watcher = new Watcher(exp, scope, function (newVal) {
      updateFn && updateFn(node, newVal, prop);
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
      node.addEventListener('input', function (e) {
        node.isInputting = true;   // 由于上面绑定了自动更新，循环依赖了，中文输入法不能用。这里加入一个标志避开自动update
        var newValue = e.target.value;
        scope[exp] = newValue;
      });
    }
  },

  // @FIXME 变更需要重新编译子元素
  htmlHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, dir);
  },

  textHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, dir);
  },

  showHandler: function (node, scope, exp, dir) {
    this.bindWatcher(node, scope, exp, 'style', 'display')
  },

  ifHandler: function (node, scope, exp, dir) {

  },

  // 属性指令 v-bind:id="id", v-bind:class="cls"
  bindHandler: function (node, scope, exp, dir) {
    var attr = dir.split(':')[1];
    switch (attr) {
      case 'class':
        // 拼成 "baseCls "+(a?"acls ":"")+(b?"bcls ":"")的形式
        exp = '"' + node.className + ' "+' + parseClassExp(exp);
        break;
      case 'style':
        // style可以使用style.cssText/node.setAttribute('style','your style')全量更新，也可以使用style.prop单个更新
        // 全量更新只需要监听全量表达式即可，但是初次编译之后其他地方脚本改了propB的话，下一次更新propA也会使用vm的值去覆盖更改后的propB
        // 单个更新的话需要监听多个值，但是不同样式之间无影响，比如初次编译后脚本更改了propC，下一次更新propB是不会影响到propC的
        // 这里使用全量更新，样式写法是这样的：<div v-bind:style="{ color: activeColor, font-size: fontSize }"></div>
        var styleStr = node.getAttribute('style');
        exp = '"' + styleStr + ';"+' + parseStyleExp(exp);
        break;
      default:

    }
    this.bindWatcher(node, scope, exp, 'attr', attr)
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
    // 当有输入的时候循环依赖了，中文输入法不能用。这里加入一个标志避开自动update
    if (!node.isInputting) {
      node.value = newVal ? newVal : '';
    }
    node.isInputting = false;  // 记得要重置标志
  },
  attr: function (node, newVal, attrName) {
    newVal = typeof newVal === 'undefined' ? '' : newVal;
    node.setAttribute(attrName, newVal);
  },
  style: function (node, newVal, attrName) {
    newVal = typeof newVal === 'undefined' ? '' : newVal;
    if (attrName === 'display') {
      newVal = newVal ? 'initial' : 'none';
    }
    node.style[attrName] = newVal;
  },
};