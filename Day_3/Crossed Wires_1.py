import os


def wire_parser(data):
    return list(map(lambda x: (x[0], int(x[1:])), data.split(',')))


def coords_computing(wire):
    x, y = 0, 0
    coords_lst = []
    for i in wire:
        if i[0] == 'R':
            lst = [(trace, y)
                   for trace in range(min(x, x+i[1]), max(x, x+i[1])+1)]
            x += i[1]
        elif i[0] == 'L':
            lst = [(trace, y)
                   for trace in range(min(x, x-i[1]), max(x, x-i[1])+1)]
            x -= i[1]

        elif i[0] == 'U':
            lst = [(x, trace)
                   for trace in range(min(y, y+i[1]), max(y, y+i[1])+1)]
            y += i[1]

        elif i[0] == 'D':
            lst = [(x, trace)
                   for trace in range(min(y, y-i[1]), max(y, y-i[1])+1)]
            y -= i[1]

        coords_lst += lst
    return set(coords_lst)


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+'\input_3.txt') as file:
        dt = file.readlines()
        f_wire = wire_parser(dt[0])
        s_wire = wire_parser(dt[1])

        f_result = coords_computing(f_wire)
        s_result = coords_computing(s_wire)
        intersec = f_result.intersection(s_result)

        nearest = [abs(i[0])+abs(i[1]) for i in intersec]
        nearest.pop(0)  # 0,0 does not count
        print(min(nearest))


if __name__ == "__main__":
    main()
