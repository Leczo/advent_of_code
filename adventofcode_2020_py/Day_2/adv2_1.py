
import os


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input2') as file:
        def parse(data):
            def inner_parse(p_data):
                return (list(map(int, p_data[0].split('-'))), p_data[1])

            x = data.strip().split(':')
            return (inner_parse(x[0].split(' ')), x[1])

        data = [parse(i) for i in file.readlines()]
        proper_passwords = sum(condition[0][0] <= string.count(condition[1]) <= condition[0][1]
                               for condition, string in data)

        print(proper_passwords)


if __name__ == "__main__":
    main()
