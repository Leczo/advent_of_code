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
    doubl_incl = [i for i in dec_excl if max(grouping(i)) >= 2]
    end = time.time()
    print(len(doubl_incl))
    print(end-start)


if __name__ == "__main__":
    main()
