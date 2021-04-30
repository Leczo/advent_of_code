import os
from itertools import permutations


def instruction_parser(i, dt):
    return list(map(int, '0'*(5-len(str(dt[i])))+str(dt[i])))


def computing(setting, inputnr, dt):
    i = 0
    temp = 0
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
            inputt = setting if temp == 0 else inputnr
            temp = temp + 1
            dt[mode(p1, 1)] = inputt
            i += 2
        elif opcode == 4:
            return dt[mode(p1, 1)]
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


def signal_counter(dt, amp_settings):
    signals = []
    for settings in amp_settings:
        counting = 0
        for setting in settings:
            counting = computing(setting, counting, dt)

        settings_str = ''.join(list(map(str, settings)))
        signals.append((counting, settings_str))
    return signals


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_7.txt') as file:
        dt = list(map(int, file.read().split(',')))
        amp_settings = list(permutations(list(range(5))))
        best_signal = max(signal_counter(dt, amp_settings), key=lambda x: x[0])
        print(best_signal)


if __name__ == "__main__":
    main()
