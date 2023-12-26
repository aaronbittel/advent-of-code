import time
from collections import namedtuple


instruction = namedtuple("instruction", ["var", "condition", "target"])


def load(file):
    with open(file, "r") as f:
        return [s.split("\n") for s in f.read().split("\n\n")]


def solve(p):
    rules, parts = p[0], p[1]
    rules_dict = create_rules(rules)
    part_numbers = create_part_numbers(parts)

    return sum(
        sum(part.values())
        for part in part_numbers
        if apply_rules(part, "in", rules_dict)
    )


def apply_rules(part_num, current_rule, rules):
    if current_rule == "R":
        return False
    elif current_rule == "A":
        return True

    start = rules[current_rule]
    instructions = start["instructions"]
    for inst in instructions:
        var, condition, target = inst
        number = part_num[var]
        if condition.get("l", None):
            check_condition = number < condition["l"]
        else:
            check_condition = number > condition["g"]

        if check_condition:
            # print(target)
            return apply_rules(part_num, target, rules)
    else:
        # print(start["destination"])
        return apply_rules(part_num, start["destination"], rules)


def create_part_numbers(parts):
    part_numbers = []
    for part in parts:
        my_dict = {}
        for x in part[1:-1].split(","):
            var, num = x.split("=")
            my_dict[var] = int(num)
        part_numbers.append(my_dict)
    return part_numbers


def create_rules(rules):
    rules_list = {}
    for r in rules:
        name, instructions = r.split("{")
        instructions = instructions.replace("}", "").split(",")
        destination, instructions = instructions[-1], instructions[:-1]

        instruction_list = []

        for inst in instructions:
            var = inst[0]
            condition, target = inst[1:].split(":")
            con = "g" if condition[0] == ">" else "l"
            con_num = int(condition[1:])
            instruct = instruction(var, {con: con_num}, target)
            instruction_list.append(instruct)

        ins = {"instructions": instruction_list, "destination": destination}

        rules_list[name] = ins
    return rules_list


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day19.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
