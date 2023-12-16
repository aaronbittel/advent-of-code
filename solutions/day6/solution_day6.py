import time
import re
import math

# starting speed: 0 mm / ms
# speed increases: 1 mm / ms


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve_part1(p):
    times, distances = [list(map(int, re.findall("\\d+", part))) for part in p]
    run = []
    for time_, distance in zip(times, distances):
        # possibilities = sum(1 for i in range(1, time_) if i * (time_ - i) > distance)

        possibilities = 0
        for i in range(1, time_):
            run_distance = i * (time_ - i)
            if run_distance > distance:
                possibilities += 1

        run.append(possibilities)
    return math.prod(run)


def solve_part2(p):
    time_, distance = [list(re.findall("\\d+", part)) for part in p]
    time_, distance = int("".join(time_)), int("".join(distance))
    possibilities = 0
    for t in range((time_ // 2) - 1, 0, -1):
        if t * (time_ - t) > distance:
            possibilities += 1
        else:
            return possibilities * 2 + 1


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution_part1 = solve_part1(load("puzzle_input_day6.txt"))
    solution_part2 = solve_part2(load("puzzle_input_day6.txt"))
    print(f"Part 1: {solution_part1, solution_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
