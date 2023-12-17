import time
import re


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    part1 = 0

    directions = [direction.replace("L", "0").replace("R", "1") for direction in p[0]]

    starting_location = "AAA"
    paths = {
        start: [left, right]
        for start, left, right in (re.findall(r"[A-Z]{3}", row) for row in p[2:])
    }

    direction_index = 0
    current_location = starting_location

    while current_location != "ZZZ":
        current_location = paths[current_location][int(directions[direction_index])]
        direction_index = (direction_index + 1) % len(directions)
        part1 += 1

    return part1


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day8.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
