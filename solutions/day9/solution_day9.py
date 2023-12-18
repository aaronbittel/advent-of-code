import time
from functools import reduce


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    part1 = 0
    histories = [list(map(int, row.split())) for row in p]
    for history in histories:
        adds = []
        diff = history
        # print(history)
        should_add = True
        while sum(diff) != 0:
            if len(diff) == 1 and diff[0] != 0:
                should_add = False
            adds.append(diff[-1])
            diff = [abs(left - right) for left, right in zip(diff, diff[1:])]
            # print(diff, "Sum: ", sum(diff))
        # print(adds)
        # print("-" * 20)
        part1 += reduce(lambda x, y: x + y, adds) if should_add else 0
    return part1


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day9.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
