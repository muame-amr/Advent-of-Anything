from collections import defaultdict


def main():
    grid = read_input("input.txt")
    # grid = read_input("test.txt")
    r_size = len(grid)
    gear_to_num = defaultdict(list)
    total = 0

    for row in range(r_size):
        num = 0
        c_size = len(grid[row])
        gears = set()  # position in array index of the gears
        for col in range(c_size + 1):  # checking if after the number is endline
            if col < c_size and grid[row][col].isdigit():
                ele = grid[row][col]
                num = num * 10 + int(ele)
                for r_adj in [-1, 0, 1]:
                    for c_adj in [-1, 0, 1]:
                        if 0 <= row + r_adj < r_size and 0 <= col + c_adj < c_size:
                            adj = grid[row + r_adj][col + c_adj]
                            if adj == "*":
                                gears.add(
                                    (row + r_adj, col + c_adj)
                                )  # store gear index position
            elif num > 0:
                for gear in gears:
                    gear_to_num[gear].append(num)
                num = 0
                gears = set()

    for k, v in gear_to_num.items():
        if len(v) == 2:
            print(f"Gear {k} : {v}")
            total += v[0] * v[1]

    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [[col for col in lines.strip()] for lines in f.readlines()]


if __name__ == "__main__":
    main()  # Answer: 80694070
