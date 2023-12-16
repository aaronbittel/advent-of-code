import time
from collections import namedtuple


mapper = namedtuple("mapper", ["destination", "source", "range"])


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f if row != "\n"]


def solve(p):
    seeds = [int(seed) for seed in p[0].split(" ")[1:]]
    sections = []
    index = -1
    for row in p[1:]:
        if not row[0].isdigit():
            sections.append([])
            index += 1
        else:
            sections[index].append(mapper(*(map(int, row.split(" ")))))

    return min(seed_conversion(seed, sections) for seed in seeds)


def seed_conversion(seed, sections):
    # print(seed, end=" -> ")
    for section in sections:
        for mapping in section:
            if mapping.source > seed:
                continue
            elif mapping.source + mapping.range < seed:
                continue
            seed += mapping.destination - mapping.source
            break
        # print(seed, end=" -> ")
    # print()
    return seed


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day5.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
