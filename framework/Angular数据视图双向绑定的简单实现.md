## 数据视图双向绑定的简单实现

先给出一个简单的例子：

HTML：
```HTML
<h3>mvvm simple model</h3>
<div>
  <p>绑定测试： <span ng-bind="mvdata"></span></p>
  <p>不同绑定测试： <span mv-bind="mvother"></span></p>
  <input type="text" mv-model="mvdata" value="" onkeyup="update(this)"/>
  <input type="text" mv-model="mvother" value="" onkeyup="update(this)"/>
  <button onclick="resetdata()">数据更新视图</button>
</div>
```
js：
```javascript
var mv = {};
var bindDom = document.querySelectorAll('[mv-bind]');

function update(item) {
  mv[item.getAttribute('mv-model')] = item.value;
  apply();
}

function apply() {
  bindDom.forEach(function (item) {
    item.innerText = mv[item.getAttribute('mv-bind')];
  });
}

function resetdata() {
  mv['mvdata'] = 'reset';
  apply();
}
```

这里分视图到数据、数据到视图两步来看。

1. 视图到数据

这里用的是监听事件的方式，框架去监听控件对应的输入事件，比如input的keydown、keyup、onchange等事件，根据用户的输入去更新mv-model绑定的变量。事件绑定也是类似，通过监听用户事件去调用被绑定的函数。这一步就是原生的addEventListener实现。

2. 数据到视图

实现数据到视图更新的是apply方法，将所有被绑定得dom节点的值更新成mv里面的值。所谓的数据自动更新，其实也就是调用了像apply一样的更新方法来实现数据“自动”更新到视图而已。

### 改进一，封装作用域

以上简单实现了一个视图-数据的双向绑定，但是这种实现太粗暴原始，离真正的工程应用还差得远。

首先是增加作用域限制。
目前的操作都是在全局作用域下进行的，范围过大容易冲突、全局检查的性能浪费也非常大，所以可以指定在局部范围内进行操作。
为了将模板和js作用域联系起来，可以在html中加入ng-scope来声明作用域范围，同时在js中使用类的方式声明Scope，代码如下：

HTML：
```HTML
<body onload="init()">
<h3>第一个scope</h3>
<div ng-scope>
  <p>绑定测试： <span ng-bind="test"></span></p>
  <p>依赖数据测试： <span ng-bind="dependence"></span></p>
  <input type="text" ng-model="test" value=""/>
  <button ng-click="resetdata()">重置数据</button>
</div>
<br>
<hr>
<h3>第二个scope</h3>
<div ng-scope>
  <p>绑定测试： <span ng-bind="test"></span></p>
  <p>依赖数据测试： <span ng-bind="dependence"></span></p>
  <input type="text" ng-model="test" value=""/>
  <button ng-click="resetdata()">重置数据</button>
</div>
</body>
```

```javascript
function Scope(domEl) {
  this.$scope = {};
  this.dom = domEl || document.body;

  var self = this;
  var bindDom = null;
  var modelDom = null;
  var clickDom = null;

  // 无依赖的数据、方法
  function declareData() {
    self.$scope.test = 'test';
    self.$scope.resetdata = function () {
      self.$scope.test = 'reset';
    };
  }

  // 有依赖的数据、方法
  function computeData() {
    self.$scope.dependence = 'dependence: ' + self.$scope.test;
  }

  function init() {
    modelDom = self.dom.querySelectorAll('[ng-model]');
    modelDom.forEach(function (el) {
      // ng接管各种用户输入事件来更新绑定数据，最后再调用apply去更新视图
      el.addEventListener('keyup', function (event) {
        self.$scope[el.getAttribute('ng-model')] = el.value;
        self.apply();
      });
      // el.addEventListener('keydown', function (e) {});
      // el.addEventListener('change', function (e) {}); ……
    });

    // 绑定用户事件
    clickDom = self.dom.querySelectorAll('[ng-click]');
    clickDom.forEach(function (el) {
      el.addEventListener('click', function (event) {
        self.$scope[el.getAttribute('ng-click').replace('()', '')]();
        self.apply();
      })
    });

    // 指定绑定数据的元素
    bindDom = self.dom.querySelectorAll('[ng-bind]');

    // 初始化数据
    declareData();

    // 应用数据并更新视图
    self.apply();
  }

  this.apply = function () {
    // 计算依赖数据
    computeData();

    // 更新视图
    bindDom.forEach(function (item) {
      item.innerText = self.$scope[item.getAttribute('ng-bind')];
    });
  }

  init();
}

function init() {
  var scopes = document.querySelectorAll('[ng-scope]');
  scopes.forEach(function (el) {
    new Scope(el);
  })
}
```

### 改进二，增加脏值检查

angular中并没有根据有无依赖来区分数据，而是使用脏检查和递归查找的方式来确认数据模型的更新。



首先是更新机制。由于DOM操作的性能消耗比较大，所以要避免现在的全量更新，应该只更新差异点。因此这里要引入一个脏值检查机制，用来筛选出mv中被更新的值再去触发视图的更新。






