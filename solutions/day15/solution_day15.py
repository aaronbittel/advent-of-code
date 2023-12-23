import time


def load(file):
    with open(file, "r") as f:
        return "".join([row.strip() for row in f]).split(",")


def solve(p):
    part1 = sum(my_hash(seq) for seq in p)
    return part1


def my_hash(string):
    res = 0
    for c in string:
        res += ord(c)
        res *= 17
        res %= 256
    return res


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day15.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
