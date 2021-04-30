
import os


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input2') as file:
        data = sorted([int(i.rstrip()) for i in file.readlines()])
        arr_len = len(data)
        for _ in range(arr_len):
            fromula = data[0] + data[1] + data[-1]
            if fromula == 2020:
                print('answer')
                print(data[0] * data[1] * data[-1])
                break
            if fromula > 2020:
                data.pop(-1)
                continue
            if fromula < 2020:
                if data[1] - data[0] < 2020 - fromula and data[2] - data[1] < 2020 - fromula:
                    data.pop(0)
                else:
                    data.pop(1)
                continue


if __name__ == "__main__":
    main()
