import re


def main():
    total = 0
    lines = read_input("input.txt")
    # lines = read_input("test.txt")
    for idx, line in enumerate(lines, start=1):
        total += calc_cubes(idx, line)

    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [lines.split(":")[1].strip() for lines in f.readlines()]


def calc_cubes(game, bag):
    for sets in bag.split(";"):
        cubes = {"green": 0, "blue": 0, "red": 0}
        sets = re.split(" |, ", sets.strip())
        size_n = len(sets)
        for i in range(1, size_n, 2):
            cubes[sets[i]] += int(sets[i - 1])
            if cubes["red"] > 12 or cubes["green"] > 13 or cubes["blue"] > 14:
                return 0
    return game


if __name__ == "__main__":
    main()  # Answer: 2176
