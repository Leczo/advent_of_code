import os


def wire_parser(data):
    return list(map(lambda x: (x[0], int(x[1:])), data.split(',')))


def coordsecond_computing(wire):
    x, y = 0, 0
    coordsecond_lst = []
    for i in wire:
        if i[0] == 'R':
            lst = [(trace, y)
                   for trace in range(min(x+1, x+i[1]+1), max(x+1, x+i[1])+1)]
            x += i[1]
        elif i[0] == 'L':
            lst = [(trace, y)
                   for trace in range(max(x-1, x-i[1]-1), min(x-1, x-i[1]-1), -1)]
            x -= i[1]

        elif i[0] == 'U':
            lst = [(x, trace)
                   for trace in range(min(y+1, y+i[1]+1), max(y+1, y+i[1])+1)]
            y += i[1]

        elif i[0] == 'D':
            lst = [(x, trace)
                   for trace in range(max(y-1, y-i[1]-1), min(y-1, y-i[1]-1), -1)]
            y -= i[1]

        coordsecond_lst += lst

    return coordsecond_lst


def step_counter(coordsecond_list, intersec):
    results = []
    for steps, coords in enumerate(coordsecond_list):
        if coords in intersec and coords != (0, 0):
            results.append([coords, steps])
    return results


def step_sort(steps):
    return sorted(steps, key=lambda x: x[0])


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_3.txt') as file:
        dt = file.readlines()
        first_wire = wire_parser(dt[0])
        second_wire = wire_parser(dt[1])

        first_wire_coords = coordsecond_computing(first_wire)
        second_wire_coords = coordsecond_computing(second_wire)
        intersec = list(
            set(first_wire_coords).intersection(set(second_wire_coords)))

        first_steps = step_sort(step_counter(first_wire_coords, intersec))
        second_steps = step_sort(step_counter(second_wire_coords, intersec))
        compare = min(i[1]+j[1] for i, j in zip(second_steps, first_steps))

        #print(first_steps, '\n', second_steps)
        print(compare+2)  # adding 2 steps for (0,0) in first and second wire


if __name__ == "__main__":
    main()
