
#include <array>
#include <condition_variable>
#include <cstdint>
#include <mutex>
#include <utility>
namespace concurrent {
template <typename T, std::size_t cap> class BlockingQueue {
public:
  BlockingQueue(){};

  ~BlockingQueue(){};

  // block while full
  void push(const T &val) {
    std::unique_lock<std::mutex> lock(mutex_);
    producer_cond_.wait(lock, [this] { return w_idx_ - r_idx_ < cap; });
    data_[w_idx_ % cap] = val;
    w_idx_++;
    consumer_cond_.notify_one();
  };

  // block while empty
  T pop() {
    std::unique_lock<std::mutex> lock(mutex_);
    consumer_cond_.wait(lock, [this] { return w_idx_ > r_idx_; });
    auto idx = r_idx_ % cap;
    r_idx_++;
    producer_cond_.notify_one();
    return data_[idx];
  };

  int size() {
    std::lock_guard<std::mutex> lock(mutex_);
    return w_idx_ - r_idx_;
  }

private:
  std::array<T, cap> data_;
  std::int64_t r_idx_ = 0;
  std::int64_t w_idx_ = 0;
  std::mutex mutex_;
  std::condition_variable producer_cond_;
  std::condition_variable consumer_cond_;
};
} // namespace concurrent