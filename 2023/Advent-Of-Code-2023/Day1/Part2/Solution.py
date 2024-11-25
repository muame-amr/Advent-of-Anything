from string import digits
import re

# Using map/dictionary
digits_mp = {
    "one": "o1e",
    "two": "t2o",
    "three": "t3e",
    "four": "4",
    "five": "5e",
    "six": "6",
    "seven": "7n",
    "eight": "e8t",
    "nine": "n9e",
}


def main():
    lines = read_input("input.txt")
    total = sum([int(extract_sum(line)) for line in lines])

    print(f"[Answer]: {total}")


def read_input(filename):
    with open(filename) as f:
        return [lines.strip() for lines in f.readlines()]


def extract_sum(original_string):
    new_string = transform_string(original_string)
    size_n = len(new_string)
    return new_string[0] + new_string[size_n - 1]


def transform_string(string_to_replace):
    for key, val in digits_mp.items():
        string_to_replace = string_to_replace.replace(key, val)
    return re.sub(r"\D", "", string_to_replace)


if __name__ == "__main__":
    main()  # Answer: 54203
