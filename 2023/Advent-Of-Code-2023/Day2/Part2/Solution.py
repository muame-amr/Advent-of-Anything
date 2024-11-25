import re
from functools import reduce
import operator


def main():
    total = 0
    lines = read_input("input.txt")
    # lines = read_input("test.txt")
    for line in lines:
        total += calc_cubes(line)

    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [lines.split(":")[1].strip() for lines in f.readlines()]


def calc_cubes(bag):
    cubes = {"green": 0, "blue": 0, "red": 0}
    for sets in bag.split(";"):
        sets = re.split(" |, ", sets.strip())
        size_n = len(sets)
        for i in range(1, size_n, 2):
            cubes[sets[i]] = max(int(sets[i - 1]), cubes[sets[i]])

    val = list(cubes.values())
    product = reduce(operator.mul, val, 1)
    print(product)

    return product


if __name__ == "__main__":
    main()  # Answer: 63700
