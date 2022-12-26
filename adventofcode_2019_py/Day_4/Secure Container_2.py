import itertools as it
import time


def grouping(arr):
    return list(map(lambda x: len(x), [list(g) for k, g in it.groupby(arr)]))


def main():
    input = [264360, 746325]
    start = time.time()
    pasw_range = list(range(input[0], input[1]))
    dec_excl = [list(str(i)) for i in pasw_range if i ==
                int(''.join(sorted(list(str(i)))))]
    just_double = [''.join(i) for i in dec_excl if 2 in grouping(i)]
    end = time.time()
    print(len(just_double))
    print(end-start)


if __name__ == "__main__":
    main()
