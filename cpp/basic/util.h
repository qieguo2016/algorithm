
#include <iostream>
#include <string>
#include <vector>
namespace basic {
static void PrintVector(const std::string &msg, const std::vector<int> &input) {
  std::cout << msg << "[";
  for (const auto &item : input) {
    std::cout << item << ",";
  }
  std::cout << "]" << std::endl;
}
} // namespace basic