
#include "sort.h"

int main(int argc, const char **argv) {
  basic::Sort s;
  std::vector<int> arr = {3, 8, 9, 1, 8, 3, 4, 2};
  basic::PrintVector("origin: ", arr);
  s.QuickSort(arr);
  basic::PrintVector("sorted: ", arr);
}