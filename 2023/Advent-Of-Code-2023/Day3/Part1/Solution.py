def main():
    grid = read_input("input.txt")
    # grid = read_input("test.txt")
    r_size = len(grid)
    total = 0

    for row in range(r_size):
        c_size = len(grid[row])
        num = 0
        has_part = False
        for col in range(c_size + 1):  # Checking if after the number is endline
            if col < c_size and grid[row][col].isdigit():
                ele = grid[row][col]
                num = num * 10 + int(ele)
                for r_adj in [-1, 0, 1]:
                    for c_adj in [-1, 0, 1]:
                        if 0 <= row + r_adj < r_size and 0 <= col + c_adj < c_size:
                            adj = grid[row + r_adj][col + c_adj]
                            if not adj.isdigit() and adj != ".":
                                has_part = True
            elif num > 0:
                if has_part:
                    print(num)
                    total += num
                has_part = False
                num = 0

        # print()

    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [[col for col in lines.strip()] for lines in f.readlines()]


if __name__ == "__main__":
    main()  # Answer: 521601
