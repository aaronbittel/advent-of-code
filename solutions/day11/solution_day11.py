import time
from itertools import combinations


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p, expanding=1):
    galaxies = {
        (x, y) for y, row in enumerate(p) for x, sym in enumerate(row) if sym == "#"
    }

    empty_rows = [i for i, row in enumerate(p) if all(sym == "." for sym in row)]
    empty_cols = [i for i in range(len(p[0])) if all(row[i] == "." for row in p)]

    galaxies = expand(galaxies, empty_rows, "row", expanding)
    galaxies = expand(galaxies, empty_cols, "col", expanding)

    part1 = sum(
        abs(x[0] - y[0]) + abs(x[1] - y[1]) for x, y in combinations(galaxies, 2)
    )

    return part1


def expand(galaxies, list_, mode, expanding):
    for item in reversed(list_):
        if mode == "row":
            galaxies = {
                (x, y + expanding) if y > item else (x, y) for (x, y) in galaxies
            }
        elif mode == "col":
            galaxies = {
                (x + expanding, y) if x > item else (x, y) for (x, y) in galaxies
            }
    return galaxies


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day11.txt")
    sol_part1 = solve(puzzle)
    sol_part2 = solve(puzzle, 1000000 - 1)
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
