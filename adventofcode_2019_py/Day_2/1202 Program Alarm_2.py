import itertools
import os


def computing(data):
    for i in range(0, len(data), 4):
        if data[i] == 99:
            break
        if data[i] == 1:
            data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
            continue
        if data[i] == 2:
            data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
            continue
        if data[i] == 3:
            continue
        if data[i] == 4:

    return data


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_2.txt') as file:
        dt = list(map(int, file.read().split(',')))

        products = list(itertools.product(
            *[list(range(0, 99)), list(range(0, 99))]))
        for tup in products:
            data, data[1], data[2] = dt[:], tup[0], tup[1]
            if computing(data)[0] == 19690720:
                print(tup)


if __name__ == "__main__":
    main()
