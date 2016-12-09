/**
 * @authors     : qieguo
 * @date        : 2016/12/09
 * @version     : 1.0
 * @description : 一个简单的观察者模式事件系统实现
 */

function EventHandle() {
    var events = {};
    this.on = function(event, callback) {
        callback = callback || function() { };
        if (typeof events[event] === 'undefined') {
            events[event] = [callback];
        } else {
            events[event].push(callback);
        }
    };

    this.emit = function(event, args) {
        if (typeof events[event] !== 'undefined') {
            events[event].forEach(function(fn) {
                fn(args);
            });
        } else {
            throw new Error('event: ' + event + ', not found');
        }
    };

    this.off = function(event) {
        if (typeof events[event] !== 'undefined') {
            delete events[event];
        }
    };
}

function test() {
    var eh = new EventHandle();

    eh.on('greet', function(str) {
        console.log(str);
    });

    eh.on('greet', function(name) {
        console.log(name + ', hello!');
    });

    eh.on('bye', function(name) {
        console.log(name + ', goodbye!');
    });

    console.log('======  start  ======');
    eh.emit('greet', 'Green');
    eh.emit('bye', 'Mark');

    console.log('======  removeListener  ======');
    eh.off('bye');
    eh.emit('bye', 'Mark');
}

test();