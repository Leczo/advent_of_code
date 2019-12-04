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
    return data


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+'\input_2.txt') as file:
        data = list(map(int, file.read().split(',')))
        data[1], data[2] = 12, 2
        print(computing(data)[0])


if __name__ == "__main__":
    main()
