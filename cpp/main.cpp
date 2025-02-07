#include <cmath>
#include <cstddef>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <iterator>
#include <random>

using namespace std;

int main() {
  random_device rd;
  mt19937 eng(rd());
  uniform_real_distribution<> distr(1, 100);

  size_t epochs{20};
  float train[][2]{
      {0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8},
  };

  float learningRate{0.01};
  float weight{(float)distr(eng)};
  float loss{0};

  for (size_t epoch = 0; epoch < epochs; ++epoch) {
    float mse{0};
    printf("Epoch: %zu, ", epoch);
    for (size_t input = 0; input < size(train); ++input) {
      float x{train[input][0]};
      float y{train[input][1]};
      float yHat{x * weight};

      loss = powf(y - yHat, 2);
      mse += loss;
      float gradient{2 * x * (yHat - y)};
      weight -= learningRate * gradient;
    }
    printf("MSE: %f\n", mse);
  }

  cout << "Weight: " << weight << endl;

  return EXIT_SUCCESS;
}
