/**
 * @authors     : qieguo
 * @date        : 2016/12/13
 * @version     : 1.0
 * @description : Compiler
 */

'use strict';

const ignorableReg = /[^\t\n\r]*/;

function Compiler(options) {
	// create node
	this.$el = typeof options.el === 'string'
		? document.querySelector(options.el)
		: options.el || document.createElement('div');

	// to documentFragment
	if (this.$el) {
		this.$fragment = nodeToFragment(this.$el);
		this.compile();
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
	return (node.nodeType == 8) || ((node.nodeType == 3) && !(ignorableReg.test(node.textContent)));
}

Compiler.prototype.compile = function (node) {
	var children = node.childNodes;
	var self = this;
	for (var i = 0, len = children.length; i < len; i++) {
		var current = children[i];
		if (current.childNodes.length >= 0) {
			self.compile(current);
		}
		if (current.nodeType === 3) {
			self.compileText(current);
		} else {
			self.compileElement(current);
		}
	}
}

Compiler.prototype.compileText = function (node) {
	console.log('compileText', node);
}

Compiler.prototype.compileElement = function (node) {
	console.log('compileElement', node);
}