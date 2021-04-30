
import os


def main():

    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input1') as file:
        data = file.readlines()
        data_to_int = list(map(int, data))
        diff_arr = [(2020 - int(i)) for i in data_to_int]

        result_array = [i for i in data_to_int if i in diff_arr]

        print(result_array[0]*result_array[1])


if __name__ == "__main__":
    main()
