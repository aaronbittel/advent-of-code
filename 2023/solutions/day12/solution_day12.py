import time
import re
from collections import namedtuple
from itertools import product


spring_record = namedtuple("spring_record", ["row", "numbers"])


def load(file):
    with open(file, "r") as f:
        return [row.strip().split() for row in f]


def solve(p):
    part1 = 0
    p: list[spring_record] = [
        spring_record(re.findall(r"[#?]+", row), list(map(int, numbers.split(","))))
        for row, numbers in p
    ]

    for i, record in enumerate(p):
        count = generate_combinations(record)
        part1 += count
        print(f"{i / len(p)} % finished.")

    return part1


def generate_combinations(record: spring_record):
    numbers = record.numbers
    string = ".".join(batch for batch in record.row)
    count_placeholder = string.count("?")
    combinations = []
    for i, combination in enumerate(product("#.", repeat=count_placeholder)):
        index = 0
        combinations.append("")
        for char in string:
            if char == "?":
                combinations[i] += combination[index]
                index += 1
            else:
                combinations[i] += char

    valid_combinations = 0

    for comb in combinations:
        batch_numbers = [len(amount) for amount in re.findall(r"[#]+", comb)]
        if batch_numbers == numbers:
            valid_combinations += 1

    return valid_combinations


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day12.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
