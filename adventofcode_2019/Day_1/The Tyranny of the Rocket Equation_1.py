import os


def main():

    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_1.txt') as file:
        data = file.readlines()
        counted_fuel = sum(map(lambda x: int(x)//3-2, data))
        print(counted_fuel)


if __name__ == "__main__":
    main()
