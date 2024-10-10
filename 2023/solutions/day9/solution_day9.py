import time
from functools import reduce
import logging

logging.basicConfig(level=logging.INFO)


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    part1, part2 = 0, 0
    histories = [list(map(int, row.split())) for row in p]

    for history in histories:
        adds = []
        subs = []
        diff = history
        logging.debug(history)
        while not all(num == 0 for num in diff):
            adds.append(diff[-1])
            subs.append(diff[0])
            diff = [right - left for left, right in zip(diff, diff[1:])]
            logging.debug(f"{diff} Sum: {sum(diff)}")
        logging.debug(f"Adds: {adds}")
        logging.debug(f"Subs: {subs}")
        logging.debug("-" * 20)
        part1 += sum(adds)
        curr = 0
        for sub in subs[::-1]:
            curr = sub - curr
        part2 += curr

    return part1, part2


if __name__ == "__main__":
    time_start = time.perf_counter()
    sol_part1, sol_part2 = solve(load("puzzle_input_day9.txt"))
    print(f"Part 1: {sol_part1}, Part2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
