import time
import shapely


RIGHT = (1, 0)
LEFT = (-1, 0)
UP = (0, -1)
DOWN = (0, 1)


direction_map = {"R": RIGHT, "L": LEFT, "U": UP, "D": DOWN}


def load(file):
    with open(file, "r") as f:
        return [row.strip().split() for row in f]


def solve(p):
    starting_pos = (0, 0)
    grid = []
    current_pos = starting_pos
    for direction, length, _ in p:
        curr_x, curr_y = current_pos
        dx, dy = direction_map[direction]
        current_pos = (curr_x + int(length) * dx, curr_y + int(length) * dy)
        grid.append(current_pos)
    polygon = shapely.Polygon(grid)
    return int(polygon.area + polygon.length / 2 + 1)


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day18.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
