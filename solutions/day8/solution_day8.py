import time


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    return p


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
