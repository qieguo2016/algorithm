/**
 * @authors     : qieguo
 * @date        : 2017/7/22
 * @description :
 */

const sorts = require('./sort');
const logger = {
  info: console.log,
  ok: console.log.bind(null, '\x1b[32m%s\x1B[39m'),
  error: console.error.bind(null, '\x1b[31m%s\x1B[39m')
}

function makeRandomArray(n) {
  let ret = [];
  for(let i = 0; i < n; i++) {
    ret.push(Math.floor(Math.random() * n * 10));  // get random number of [0, 10*n)
  }
  logger.info('random array: ' + JSON.stringify(ret));
  return ret;
}

function isArraySort(target) {
  for(let i = 0; i < target.length - 1; i++) {
    if(target[i] > target[i + 1]) {
      logger.error('fail: ' + JSON.stringify(target));
      return false;
    }
  }
  logger.ok('success: ' + JSON.stringify(target));
  return true;
}

function compareArray(a, b) {
  for(var i = 0; i < a.length; i++) {
    if(a[i] !== b[i]) {
      return false;
    }
  }
  return true;
}

// 以冒泡排序为基准测试
function main() {
  const origin = makeRandomArray(30);
  let sorted;
  for(let name in sorts) {
    logger.info(`======== ${name} ========`);
    if(!sorted) {
      sorted = sorts[name](origin.slice(0));
      isArraySort(sorted);
    } else {
      const s = sorts[name](origin.slice(0));
      let isOk;
      if(name === 'topSortViaHeap') {
        isOk = compareArray(sorted.slice(0, 10), s);
      } else {
        isOk = compareArray(sorted, s);
      }
      isOk ? logger.ok('success: ' + JSON.stringify(s))
        : logger.error('fail: ' + JSON.stringify(s))
    }
  }
}

main();