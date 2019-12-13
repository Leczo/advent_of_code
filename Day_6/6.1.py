import os
from itertools import chain


def root_finder(graph):
    keys = graph.keys()
    values = list(chain(*graph.values()))
    root = [i for i in keys if i not in values]
    return root


def counter(graph, root):
    steps = 0
    pointer = [root]
    branch = []
    #trace = []
    nodes_steps = {}
    while True:
        # print(pointer)
        if len(pointer) == 2:
            branch.append([pointer[1], steps])

        try:
            pointer = graph[pointer[0]]

        except:
            KeyError
            pass
            try:
                # trace.append(steps)
                pointer = [branch[0][0]]
                steps = branch[0][1] - 1
                del branch[0]
            except:
                IndexError
                break

        steps += 1
        nodes_steps.update({pointer[0]: steps})
    return nodes_steps


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_6.txt') as file:
        dt = [i.replace('\n', '').split(')') for i in file.readlines()]

        graph = {}
        for i in dt:
            if i[0] in graph.keys():
                graph[i[0]].append(i[1])
            else:
                graph.update({i[0]: [i[1]]})

        g_root = root_finder(graph)[0]
        f_branch = counter(graph, g_root)
        x = sum(i for i in f_branch.values())
        print(x)


if __name__ == "__main__":
    main()
