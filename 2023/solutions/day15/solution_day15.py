import time
from collections import defaultdict


def load(file):
    with open(file, "r") as f:
        return f.read().split(",")


def solve(p):
    part1 = sum(my_hash(seq) for seq in p)
    part2 = solve_part2(p)

    return part1, part2


def solve_part2(p):
    boxes = defaultdict(dict)
    for instruction in p:
        if "-" in instruction:
            label = instruction[:-1]
            boxes[my_hash(label)].pop(label, None)
        else:
            label, focal_length = instruction.split("=")
            boxes[my_hash(label)][label] = int(focal_length)

    return sum(
        (index + 1) * fl_nr * fl
        for index, box in boxes.items()
        for fl_nr, fl in enumerate(box.values(), start=1)
    )


def my_hash(string):
    res = 0
    for c in string:
        res = ((res + ord(c)) * 17) % 256
    return res


if __name__ == "__main__":
    time_start = time.perf_counter()
    sol_part1, sol_part2 = solve(load("puzzle_input_day15.txt"))
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
