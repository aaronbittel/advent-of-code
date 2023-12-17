import time
import re
import math


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve_part1(p):
    part1 = 0
    directions = [direction.replace("L", "0").replace("R", "1") for direction in p[0]]
    starting_location = "AAA"
    paths = {
        start: [left, right]
        for start, left, right in (re.findall(r"[A-Z]{3}", row) for row in p[2:])
    }
    direction_index = 0
    current_location = starting_location
    while current_location != "ZZZ":
        current_location = paths[current_location][int(directions[direction_index])]
        direction_index = (direction_index + 1) % len(directions)
        part1 += 1

    return part1


def solve_part2(p, starting):
    part1 = 0

    directions = [direction.replace("L", "0").replace("R", "1") for direction in p[0]]

    # starting_location = "AAA"
    starting_location = starting
    paths = {
        start: [left, right]
        for start, left, right in (re.findall(r"[A-Z]{3}", row) for row in p[2:])
    }

    direction_index = 0
    current_location = starting_location

    points = []
    times = 0
    while times <= 1:
        current_location = paths[current_location][int(directions[direction_index])]
        if current_location[-1] == "Z":
            times += 1
            points.append(part1)
        direction_index = (direction_index + 1) % len(directions)
        part1 += 1

    return points


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day8.txt")
    my_list = []
    for starting_pos in ["SLA", "AAA", "LVA", "NPA", "GDA", "RCA"]:
        solution_part1 = solve_part2(puzzle, starting_pos)
        my_list.append(solution_part1[1] - solution_part1[0])
    print(f"Part 1: {solve_part1(puzzle)}, Part 2: {math.lcm(*my_list)}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
