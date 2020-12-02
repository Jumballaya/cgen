#include <iostream>
#include "../inc/Vector.h"


int main() {

  Vector<int> v = 10;

  for (int i = 0; i < v.size(); i++) {
    v[i] = i;
  }

  for (int i = 0; i < v.size(); i++) {
    std::cout << v[i] << "\n";
  }

  return(0);
}