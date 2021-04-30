import os
from itertools import chain
from collections import OrderedDict


def root_finder(graph):
    keys = graph.keys()
    values = list(chain(*graph.values()))
    root = [i for i in keys if i not in values]
    return root


def counter(graph, root):
    steps = 0
    pointer = [root]
    branch = []
    # trace = []
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


def predecesor(node, tree_steps, graph):
    # node_value = tree_steps[node]
    keys = graph.keys()
    return [key for key, val in tree_steps.items() if key in keys and node in graph[key]].pop()


def path_counter(node, tree_steps, graph):
    pred = node
    path = []
    while tree_steps[pred] > 1:
        pred = predecesor(pred, tree_steps, graph)
        path.append(pred)
    return path


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
        tree_steps = counter(graph, g_root)
        node1, node2 = 'YOU', 'SAN'
        path1 = path_counter(node1, tree_steps, graph)
        path2 = path_counter(node2, tree_steps, graph)
        common = [i for i in path1 if i in path2][0]
        x = len(path1[:path1.index(common)])
        y = len(path2[:path2.index(common)])
        print(x+y)


if __name__ == "__main__":
    main()
