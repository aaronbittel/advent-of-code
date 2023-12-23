import time


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    part1 = 0
    row_length = len(p[0])
    p = ["".join(list(row)).split("#") for row in list(zip(*p))]
    for row in p:
        index = 0
        for part in row:
            if "O" not in part:
                index += len(part) + 1
                continue
            part1 += sum(row_length - index - x for x in range(part.count("O")))
            index += len(part) + 1
    return part1


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day14.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
