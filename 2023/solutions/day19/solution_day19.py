import time
import math
from collections import defaultdict


def load(file):
    with open(file, "r") as f:
        return f.read()


def parsing(p):
    workflows, ratings = [row.split("\n") for row in p.split("\n\n")]
    workflows2 = defaultdict(list)
    for row in workflows:
        name, conditions = row.split("{")
        for condition in conditions.split(","):
            if condition[-1] == "}":
                con = (None, condition[:-1])
            else:
                con, target = condition.split(":")
                var, op, num = con[0], con[1], int(con[2:])
                con = ((var, op, num), target)
            workflows2[name].append(con)

    ratings = [eval(f"dict({rating[1:-1]})") for rating in ratings]
    return workflows2, ratings


def solve(p):
    workflows, ratings = parsing(p)

    part1 = 0
    for rating in ratings:
        curr = "in"
        while True:
            for rule in workflows[curr]:
                con, target = rule
                if not con:
                    curr = target
                    break
                var, op, num = con
                if eval(f"{rating[var]}{op}{num}"):
                    curr = target
                    break

            if curr == "A":
                part1 += sum(rating.values())
                break

            if curr == "R":
                break

    ranges = {c: (1, 4000) for c in "xmas"}
    return part1, comb(workflows, ranges)


def comb(wfs, ranges, curr_wf="in"):
    if curr_wf == "A":
        return math.prod((high - low + 1) for low, high in ranges.values())
    if curr_wf == "R":
        return 0

    total = 0
    for con, target in wfs[curr_wf]:
        if not con:
            total += comb(wfs, ranges, target)
        else:
            var, op, num = con
            new_ranges = dict(ranges)
            low, high = ranges[var]
            if op == "<":
                new_ranges[var] = (low, num - 1)
                ranges[var] = (num, high)
            else:
                new_ranges[var] = (num + 1, high)
                ranges[var] = (low, num)
            total += comb(wfs, new_ranges, target)
    return total


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day19.txt"))
    print(f"Solution Part 1 & 2: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
