import time
from functools import reduce
import logging

logging.basicConfig(level=logging.DEBUG)


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    part1 = 0
    histories = [list(map(int, row.split())) for row in p]

    for history in histories:
        adds = []
        diff = history
        logging.debug(history)
        while not all(num == 0 for num in diff):
            adds.append(diff[-1])
            diff = [right - left for left, right in zip(diff, diff[1:])]
            logging.debug(f"{diff} Sum: {sum(diff)}")
        logging.debug(adds)
        logging.debug("-" * 20)
        part1 += sum(adds)

    return part1


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day9.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
