/**
 * @authors     : qieguo
 * @date        : 2016/12/09
 * @version     : 1.0
 * @description : 一个简单的观察者模式事件系统实现
 */

const Noop = () => {};

class EventEmitter {
    constructor() {
        this.events = [];
    }

    on(event, handler = Noop) {
        if (typeof this.events[event] === 'undefined') {
            this.events[event] = [handler];
        } else {
            this.events[event].push(handler);
        }
    }

    once(event, handler = Noop) {
        const once = `once_${event}`;
        if (typeof this.events[once] === 'undefined') {
            this.events[once] = [handler];
        } else {
            this.events[once].push(handler);
        }
    }

    emit(event, args) {
        const once = `once_${event}`;
        if (typeof this.events[once] !== 'undefined') {
            this.events[once].forEach((handler) => {
                handler(args);
            });
            delete this.events[`once_${event}`];
        }
        if (typeof this.events[event] !== 'undefined') {
            this.events[event].forEach((handler) => {
                handler(args);
            });
        }
    }

    off(event, handler) {
        if (typeof this.events[event] !== 'undefined') {
            if (!!handler) {
                const index = this.events[event].indexOf(handler);
                this.events[event].splice(index, 1);
            } else {
                delete this.events[event];
            }
        }
    };
}

function test() {
    const eh = new EventEmitter();

    const firstHd = (str) => {console.log('first greet: ', str);};

    eh.on('greet', firstHd);

    eh.on('greet', (str) => {console.log('second greet: ', str);});

    eh.on('bye', (name) => {console.log(name + ', goodbye!');});

    eh.once('break', (str) => {console.log(`once break: ${str}`);});
    eh.on('break', (str) => {console.log(`on break: ${str}`);});

    console.log('======  start  ======');
    eh.emit('greet', 'Green');
    eh.emit('bye', 'Mark');
    eh.emit('break', 'Jack');
    eh.emit('break', 'Tony');

    console.log('======  removeListener  ======');
    eh.off('bye');
    eh.emit('bye', 'Mark');

    eh.off('greet', firstHd);
    eh.emit('greet', 'Green');
}

test();