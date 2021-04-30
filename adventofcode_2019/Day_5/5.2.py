
import os


def instruction_parser(i, dt):
    return list(map(int, '0'*(5-len(str(dt[i])))+str(dt[i])))


def computing(dt):
    i = 0
    inputt = 5
    while i <= len(dt):

        def mode(parameter, number): return (
            dt[i+number], i+number)[parameter]

        p3, p2, p1, *opcode = instruction_parser(i, dt)
        opcode = int(''.join(map(str, opcode)))

        if opcode == 99:
            break
        elif opcode == 1:
            dt[mode(p3, 3)] = dt[mode(p2, 2)] + dt[mode(p1, 1)]
            i += 4
        elif opcode == 2:
            dt[mode(p3, 3)] = dt[mode(p2, 2)] * dt[mode(p1, 1)]
            i += 4
        elif opcode == 3:
            dt[mode(p1, 1)] = inputt
            i += 2
        elif opcode == 4:
            print(dt[mode(p1, 1)])
            break
        elif opcode == 5:
            i = dt[mode(p2, 2)] if dt[mode(p1, 1)] != 0 else i+3
        elif opcode == 6:
            i = dt[mode(p2, 2)] if dt[mode(p1, 1)] == 0 else i+3
        elif opcode == 7:
            dt[mode(p3, 3)] = 1 if dt[mode(p1, 1)] < dt[mode(p2, 2)] else 0
            i += 4
        elif opcode == 8:
            dt[mode(p3, 3)] = 1 if dt[mode(p1, 1)] == dt[mode(p2, 2)] else 0
            i += 4


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_5.txt') as file:
        dt = list(map(int, file.read().split(',')))
        computing(dt)


if __name__ == "__main__":
    main()
