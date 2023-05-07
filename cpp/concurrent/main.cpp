
#include "blocking_queue.h"
#include <chrono>
#include <cstdio>
#include <iostream>
#include <thread>

int main() {
  constexpr int num = 2;
  std::thread consumers[num], producers[num];
  concurrent::BlockingQueue<std::int32_t, 5> bq;
  for (int i = 0; i < num; ++i) {
    consumers[i] = std::thread([i, &bq] {
      int n = bq.pop();
      while (n > 0) {
        std::printf("---- cc (%d) ---- : %d \n", i, n);
        n = bq.pop();
      }
    });
    producers[i] = std::thread([i, &bq] {
      int n = 1;
      while (n < 10) {
        std::printf("---- pp (%d) ---- : %d \n", i, i * 100 + n);
        bq.push(i * 100 + n);
        n++;
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
      }
      bq.push(-1);
    });
  }

  // join them back:
  for (int i = 0; i < num; ++i) {
    producers[i].join();
    consumers[i].join();
  }

  return 0;
}
