import time


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def calculate_roll(p):
    return transpose(roll(transpose(p)))


def solve(p):
    part1 = calculate_load(calculate_roll(p))
    pattern = [p]
    while True:
        p = cycle(p)
        if p in pattern:
            break
        pattern.append(p)
    offset = pattern.index(p)
    cycle_length = len(pattern) - offset
    part2 = calculate_load(pattern[(1_000_000_000 - offset) % cycle_length + offset])
    return part1, part2


def cycle(p):
    for _ in range(4):
        p = transpose(p)
        p = roll(p)
        p = [row[::-1] for row in p]
    return p


def transpose(p):
    return ["".join(col) for col in zip(*p)]


def calculate_load(p: list[str]):
    return sum(row.count("O") * (len(p) - i) for i, row in enumerate(p))


def roll(p):
    return [
        "#".join("".join(sorted(sub, reverse=True)) for sub in row.split("#"))
        for row in p
    ]


if __name__ == "__main__":
    time_start = time.perf_counter()
    sol_part1, sol_part2 = solve(load("puzzle_input_day14.txt"))
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
