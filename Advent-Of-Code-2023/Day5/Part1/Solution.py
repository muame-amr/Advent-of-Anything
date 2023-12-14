import sys
import re
from collections import defaultdict


def get_maps(data):
    maps = []
    current_map = []
    for line in data[1:]:
        if line == "":
            if current_map:
                maps.append(current_map)
                current_map = []
        else:
            current_map.append(line)

    if current_map:
        maps.append(current_map)

    return maps


def get_value_based_on_source_seed(source, maps):
    source_value = source
    for m in maps:
        values = m[1:]

        for i in values:
            i = [int(x) for x in i.split(" ")]
            source_interval = range(i[1], i[1] + i[2])

            if source_value in source_interval:
                distance = source_value - i[1]
                source_value = i[0] + distance
                break
            else:
                pass

    return source_value


def main(filename):
    input_file = open(filename).read().strip()
    data = input_file.splitlines()

    seeds = [int(x) for x in data[0].split(": ")[1].split(" ")]
    maps = get_maps(data)

    final_values = [get_value_based_on_source_seed(seed, maps) for seed in seeds]
    print("[Answer]: ", min(final_values))


if __name__ == "__main__":
    main(sys.argv[1])  # Answer: 26273516
