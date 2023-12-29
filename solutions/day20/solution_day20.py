import time
import logging
import math
from enum import Enum
from abc import ABC


"""
part2:
    goal: rx low pulse
    -> dt (only connection, Conjunction Module) needs all modules connected to dt to last sent a HIGH
"""

logging.basicConfig(level=logging.INFO)


class Pulse(Enum):
    LOW = False
    HIGH = True


class Module(ABC):
    def __init__(self, name: str, connections: list[str]):
        self.name = name
        self.connections = connections

    def __repr__(self):
        return f"{self.__class__.__name__}: {self.connections}"


class Broadcaster(Module):
    def __init__(self, connections: list[str]):
        super().__init__("broadcaster", connections)


class FlipFlop(Module):
    def __init__(self, name: str, status: bool, connections: list[str]):
        super().__init__(name, connections)
        self.status = status

    def flip(self):
        self.status = True if not self.status else False

    def __repr__(self):
        return f"{super().__repr__()}, {self.status}"


class Conjunction(Module):
    def __init__(self, name: str, connections: list[str]):
        super().__init__(name, connections)
        self.connection_history = {}

    def __repr__(self):
        return f"{super().__repr__()}, {self.connection_history}"


class Untyped(Module):
    def __init__(self, name: str):
        super().__init__(name, [])


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def parse(p):
    config = {}
    for row in p:
        row_split = row.split()
        connections = [conn.replace(",", "") for conn in row_split[2:]]
        module_type = row_split[0][0]
        name = row_split[0][1:]
        match module_type:
            case "b":
                config["broadcaster"] = Broadcaster(connections)
            case "%":
                config[name] = FlipFlop(name, False, connections)
            case "&":
                config[name] = Conjunction(name, connections)

    untyped_module_names = []
    for name, module in config.items():
        for conn in module.connections:
            if conn in config:
                continue
            untyped_module_names.append(conn)
        if not isinstance(module, Conjunction):
            continue
        curr_name = module.name
        for name2, module2 in config.items():
            if curr_name in module2.connections:
                config[curr_name].connection_history[name2] = Pulse.LOW  #

    for untyped_module in untyped_module_names:
        config[untyped_module] = Untyped(untyped_module)

    return config


def iteration(config, part2, it):
    low_pulses, high_pulses = 0, 0
    q = [(config["broadcaster"], Pulse.LOW)]
    conjunction_queue = []
    logging.debug("\n")
    logging.debug(f"button -{Pulse.LOW}-> broadcaster")
    while q:
        module, pulse = q.pop(0)
        if pulse == Pulse.LOW:
            low_pulses += 1
        else:
            high_pulses += 1
        if isinstance(module, Broadcaster):
            for conn in module.connections:
                q.append((config[conn], pulse))
                logging.debug(f"{module.name} -{pulse}-> {conn}")
        elif isinstance(module, FlipFlop):
            if pulse == Pulse.HIGH:
                continue
            config[module.name].flip()
            for conn in module.connections:
                if config[module.name].status:
                    q.append((config[conn], Pulse.HIGH))
                    logging.debug(f"{module.name} -{Pulse.HIGH}-> {conn}")
                else:
                    q.append((config[conn], Pulse.LOW))
                    logging.debug(f"{module.name} -{Pulse.LOW}-> {conn}")
        elif isinstance(module, Conjunction):
            item = conjunction_queue.pop(0)
            if module.name == "dt" and pulse == Pulse.HIGH and part2[item] == 0:
                part2[item] = it
            module.connection_history[item] = pulse
            if all(pulse == Pulse.HIGH for pulse in module.connection_history.values()):
                for conn in module.connections:
                    q.append((config[conn], Pulse.LOW))
                    logging.debug(f"{module.name} -{Pulse.LOW}-> {conn}")
            else:
                for conn in module.connections:
                    if conn in config:
                        q.append((config[conn], Pulse.HIGH))
                    logging.debug(f"{module.name} -{Pulse.HIGH}-> {conn}")
        elif isinstance(module, Untyped):
            continue

        for conn in module.connections:
            if isinstance(config[conn], Conjunction):
                conjunction_queue.append(module.name)

    if not part2:
        return low_pulses, high_pulses


def solve(p, part2):
    config = parse(p)
    if not part2:
        low_pulses_count = high_pulses_count = 0
        for _ in range(1):
            low_pulses, high_pulses = iteration(config, {}, 0)
            low_pulses_count += low_pulses
            high_pulses_count += high_pulses
        return low_pulses_count * high_pulses_count
    else:
        index = 1
        part2 = {name: 0 for name in ["ks", "pm", "dl", "vk"]}
        while any(val == 0 for val in part2.values()):
            iteration(config, part2, index)
            index += 1
        return math.lcm(*(part2.values()))


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day20.txt")
    # sol_part1 = solve(puzzle, False)
    sol_part2 = solve(puzzle, True)
    # print(f"Part 1: {sol_part1}")
    print(f"Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
