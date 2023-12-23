import time


def load(file):
    with open(file, "r") as f:
        return [s.split("\n") for s in f.read().split("\n\n")]


def solve(p, part1):
    solution = 1410
    for pattern in p:
        solution += mirror(pattern, part1) * 100 + mirror(list(zip(*pattern)), part1)
    return solution


def mirror(pattern, part1):
    for axis in range(1, len(pattern)):
        above, below = reversed(pattern[:axis]), pattern[axis:]
        if part1:
            if all(a == b for a, b in zip(above, below)):
                return axis
        else:
            if sum(c1 != c2 for a, b in zip(above, below) for c1, c2 in zip(a, b)) == 1:
                return axis
    return 0


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day13.txt")
    sol_part1 = solve(puzzle, True)
    sol_part2 = solve(puzzle, False)
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
