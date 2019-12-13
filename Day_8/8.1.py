import os


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_8.txt') as file:
        dt = file.read()
        lst = [list(dt[i:i+150]) for i in range(0, len(dt), 150)]
        least_zeros = list(map(int, min(lst, key=lambda x: x.count('0'))))
        product = least_zeros.count(1)*least_zeros.count(2)
        print(product)


if __name__ == "__main__":
    main()
