def main():
    # total = 0
    lines = read_input("input.txt")
    # lines = read_input("test.txt")

    game_n = len(lines)
    multipliers = [1 for game in range(0, game_n)]

    # for line in lines:
    #     total += check_winning(line)

    for idx, line in enumerate(lines, start=1):
        check_piles(idx, line, multipliers)

    total = sum(multipliers)
    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [lines.split(":")[1].strip() for lines in f.readlines()]


def check_winning(card):
    winning_numbers, our_numbers = split_cards(card)
    size_n = sum([num in winning_numbers for num in our_numbers])
    points = pow(2, size_n - 1) if size_n > 0 else 0

    return points


def check_piles(curr_idx, card, multipliers):
    curr_mult = int(multipliers[curr_idx - 1])
    winning_numbers, our_numbers = split_cards(card)
    copy_range = sum([num in winning_numbers for num in our_numbers])

    for i in range(curr_idx, curr_idx + copy_range):
        multipliers[i] += curr_mult

    return multipliers


def split_cards(card):
    winning_numbers = card.split("|")[0].strip().split()
    our_numbers = card.split("|")[1].strip().split()

    return winning_numbers, our_numbers


if __name__ == "__main__":
    main()  # Answer: 10425665
