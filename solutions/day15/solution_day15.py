import time
from collections import namedtuple


lens = namedtuple("lens", ["label", "focal_length"])


def load(file):
    with open(file, "r") as f:
        return f.read().split(",")


def solve(p):
    part1 = sum(my_hash(seq) for seq in p)
    part2 = solve_part2(p)

    return part1, part2


def solve_part2(p):
    boxes = [[] for _ in range(len(p))]
    for instruction in p:
        instruction = instruction.replace("\n", "")
        label = "".join(c for c in instruction if c.isalpha())
        box_nr = my_hash(label)

        if "-" in instruction:
            if any(l.label == label for l in boxes[box_nr]):
                index = [i for i, l in enumerate(boxes[box_nr]) if l.label == label][0]
                boxes[box_nr].pop(index)
        else:
            lens_instance = lens(label, int(instruction[-1]))
            if any(l.label == lens_instance.label for l in boxes[box_nr]):
                index = [
                    i
                    for i, l in enumerate(boxes[box_nr])
                    if l.label == lens_instance.label
                ][0]
                boxes[box_nr].pop(index)
                boxes[box_nr].insert(index, lens_instance)
            else:
                boxes[box_nr].append(lens_instance)
    part2 = 0
    for i, box in enumerate(boxes, start=1):
        for j, l in enumerate(box, start=1):
            part2 += i * j * l.focal_length

    return part2


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
