import random

train = [
    [0, 0],
    [1, 2],
    [2, 4],
    [3, 6],
    [4, 8],
]

trainingRate = 0.1
weight = random.random() * 100.0

for epoch in range(10):
    mean = 0.0
    for input in train:
        x = input[0]
        y = input[1]
        yHat = x * weight

        loss = (y - yHat) ** 2
        mean += loss
        gradient = 2 * x * (yHat - y)
        weight -= trainingRate * gradient

    print(f"epoch:{epoch}\tweight:{abs(weight)}\tmse:{abs(mean)}")
    if abs(mean) < 0.0001:
        break


