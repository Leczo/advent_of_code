import os


def fuel_count(mass, counter=0):
    fuel = mass//3-2
    if fuel <= 0:
        return counter
    return fuel_count(fuel, counter+fuel)


def main():

    with open(os.path.dirname(os.path.realpath(__file__))+'\input_1.txt') as file:
        data = file.readlines()
        counted_fuel = sum(fuel_count(int(i)) for i in data)
        print(counted_fuel)


if __name__ == "__main__":
    main()
