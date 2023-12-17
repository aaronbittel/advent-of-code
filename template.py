import time
import re

# starting speed: 0 mm / ms
# speed increases: 1 mm / ms


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve_part1(p):
    # don't need to convert puzzle to list; can work with map object
    times, distances = [map(int, re.findall("\\d+", part)) for part in p]
    # no need to create a list, append items to list and mult all items in list, just use one variable
    part1 = 1
    # record better name than distance
    for time_, record in zip(times, distances):
        # possibilities = sum(1 for i in range(1, time_) if i * (time_ - i) > distance)

        possibilities = 0
        for i in range(1, time_):
            if i * (time_ - i) > record:
                possibilities += 1

        part1 *= possibilities
    return part1


def solve_part2(p):
    # x * (a - x) == -x^2 + ax (parabola with maximum at a // 2)
    # go from max to -1 to left side while y > distance: increase possibilities
    # return possibilities * 2 (because of symmetry) and add 1 for max

    times, records = [map(int, re.findall(r"\d+", part.replace(" ", ""))) for part in p]

    possibilities = 0
    for time_, record in zip(times, records):
        for t in range((time_ // 2) - 1, 0, -1):
            if t * (time_ - t) > record:
                possibilities += 1
            else:
                return possibilities * 2 + 1


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day6.txt")
    solution_part1 = solve_part1(puzzle)
    solution_part2 = solve_part2(puzzle)
    print(f"Part 1: {solution_part1, solution_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
