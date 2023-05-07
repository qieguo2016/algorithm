#include "util.h"
#include <vector>
namespace basic {

class Sort {
private:
  void quickSortImp(std::vector<int> &arr, int left, int right) {
    if (left >= right) {
      return;
    }
    int datum = arr[left];
    int i = left, j = right;
    while (i < j) { // 每次循环交换两个数
      // 因为arr[left]已经缓存到datum，先找右侧比基准小的，直接将arr[left]换掉
      while (i < j && arr[j] >= datum) {
        j--;
      }
      arr[i] = arr[j];
      while (i < j && arr[i] < datum) {
        i++;
      }
      arr[j] = arr[i];
    }
    arr[i] = datum;
    quickSortImp(arr, left, i - 1);
    quickSortImp(arr, i + 1, right);
  }

public:
  void QuickSort(std::vector<int> &arr) {
    return quickSortImp(arr, 0, arr.size() - 1);
  }
};

} // namespace basic