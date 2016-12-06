/**
 * @authors     : qieguo
 * @date        : 2016/11/29
 * @version     : 1.0
 * @description : js模拟链表
 */

'use strict';
function Chain(key, value) {
  this.next = null;
  this.key = key;
  this.value = value;
  this.length = 1;
}

// 插入元素
Chain.prototype.insertAfter = function (pos, key, value) {
  var currentObj = this;
  var addObj = {
    key: key,
    value: value
  }

  // 循环
  while (currentObj.key !== pos) {
    currentObj = currentObj.next;
  }

  // 找到元素
  addObj.next = currentObj.next;
  currentObj.next = addObj;
  this.length++;
  return this;
};

// 删除元素
Chain.prototype.delele = function (key) {
  var last = null;
  var currentObj = this;
  // 循环
  while (currentObj.key !== key) {
    last = currentObj;
    currentObj = currentObj.next;
  }
  // 找到元素
  last.next = currentObj.next;
  this.length--;
  return this;
};

// 查找元素
Chain.prototype.find = function (key) {
  var currentObj = this;
  // 循环
  while (currentObj.key !== key) {
    currentObj = currentObj.next;
  }
  // 找到元素
  return currentObj.value;
};

Chain.prototype.forEach = function (fn) {
  var currentObj = this;
  // 循环
  while (currentObj.next !== null) {
    fn({key: currentObj.key, value: currentObj.value});
    currentObj = currentObj.next;
  }
  fn({key: currentObj.key, value: currentObj.value});
};

function test() {
  var chain = new Chain('header', 'head value');
  chain.insertAfter('header', 'second', 'next to header');
  chain.insertAfter('second', '3rd', '3rd value');
  chain.insertAfter('3rd', '4th', '4th value');
  chain.insertAfter('4th', '5th', '5th value');
  console.log('add to end', JSON.stringify(chain));

  chain.insertAfter('3rd', 'add4th', 'add 4th value');
  console.log('insert between', JSON.stringify(chain));

  chain.delele('3rd');
  console.log('delete 3rd', JSON.stringify(chain));

  var temp = chain.delele('add4th').find('5th');
  console.log('5th: ', temp);

  console.log('chain.length', chain.length);

  chain.forEach(function (obj) {
    console.log(obj);
  });

}

test();