import os


def main():
    with open(os.path.dirname(os.path.realpath(__file__))+os.path.sep+'input_8.txt') as file:
        dt = file.read()
        lst = [list(dt[i:i+150]) for i in range(0, len(dt), 150)]
        image = [None]*len(lst[0])  # 25x6
        for layer in lst:
            for position, pixel in enumerate(layer):
                if image[position] != None or pixel == '2':
                    continue
                elif pixel == '0':
                    image[position] = ' '
                elif pixel == '1':
                    image[position] = '▪'
        image_format = [''.join(image[i:i+25]) for i in range(0, 150, 25)]
        for i in image_format:
            print(i)


if __name__ == "__main__":
    main()
