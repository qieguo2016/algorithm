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
function Scope(node, declareData, computeData) {
  this.$scope = {};
  this.dom = node || document.body;
  this.bindDom = null;

  var self = this;

  this.declareData = function () {
    declareData.call(self);
  };

  this.computeData = function () {
    computeData.call(self);
  }

  this.init();
}

Scope.prototype.apply = function () {
  var self = this;
  // 计算依赖数据
  this.computeData();
  // 更新视图
  this.bindDom.forEach(function (item) {
    item.innerText = self.$scope[item.getAttribute('ng-bind')];
  });
}

Scope.prototype.init = function () {
  var self = this;
  // watch过程，未做脏检查
  var modelDom = self.dom.querySelectorAll('[ng-model]');
  modelDom.forEach(function (el) {
    // ng接管各种用户输入事件来更新绑定数据，然后计算整个模型数据，最后再调用apply去更新视图
    el.addEventListener('keyup', function () {
      self.$scope[el.getAttribute('ng-model')] = el.value;
      self.apply();
    });
    // el.addEventListener('keydown', function (e) {});
    // el.addEventListener('change', function (e) {}); ……
  });

  // 绑定用户事件
  var clickDom = self.dom.querySelectorAll('[ng-click]');
  clickDom.forEach(function (el) {
    el.addEventListener('click', function (event) {
      self.$scope[el.getAttribute('ng-click').replace('()', '')]();
      self.apply();
    })
  });

  // 指定绑定数据的元素
  self.bindDom = self.dom.querySelectorAll('[ng-bind]');

  // 初始化数据
  self.declareData();

  // 应用数据并更新视图
  self.apply();
}

function init() {
  var scopes = document.querySelectorAll('[ng-scope]');
  new Scope(scopes[0], function () {
    var self = this;
    self.$scope.test = 'test';
    self.$scope.resetdata = function () {
      self.$scope.test = 'reset';
    };
  }, function () {
    var self = this;
    self.$scope.dependence = 'dependence: ' + self.$scope.test;
  });

  new Scope(scopes[1], function () {
    var self = this;
    self.$scope.test = '测试';
    self.$scope.resetdata = function () {
      self.$scope.test = '重置';
    };
  }, function () {
    var self = this;
    self.$scope.dependence = '依赖: ' + self.$scope.test;
  });

  //        new Scope(scopes[0]);

}
```

### 改进二，增加脏值检查

angular中并没有根据有无依赖来区分数据，而是使用脏检查和递归查找的方式来确认数据模型的更新，这一点跟Vue、knockout之类的实现是不一样的。

上述例子中的实现也是区分了简单数据和依赖数据，这样的好处是数据更新起来更加简单，但是需要框架使用者自己区分依赖数据和简单数据并区别对待，增加了一点点使用者的复杂度。

回到Angular中，首先是更新机制。由于DOM操作的性能消耗比较大，所以要避免现在的全量更新，应该只更新差异点。因此这里要引入一个脏值检查机制，用来筛选出mv中被更新的值再去触发视图的更新。






