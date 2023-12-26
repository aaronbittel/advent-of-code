import time
import shapely


RIGHT = (1, 0)
LEFT = (-1, 0)
UP = (0, -1)
DOWN = (0, 1)


direction_map = {"R": RIGHT, "L": LEFT, "U": UP, "D": DOWN}
direction_map2 = {"0": RIGHT, "1": DOWN, "2": LEFT, "3": UP}


def load(file):
    with open(file, "r") as f:
        return [row.strip().split() for row in f]


def solve(p, part1):
    starting_pos = (0, 0)
    grid = []
    current_pos = starting_pos
    for direction, length, color in p:
        if not part1:
            length = int(color[2:-2], 16)
            dx, dy = direction_map2[color[-2]]
        else:
            length = int(length)
            dx, dy = direction_map[direction]

        curr_x, curr_y = current_pos
        current_pos = (curr_x + length * dx, curr_y + length * dy)
        grid.append(current_pos)
    polygon = shapely.Polygon(grid)
    return int(polygon.area + polygon.length / 2 + 1)


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day18.txt")
    sol_part1 = solve(puzzle, True)
    sol_part2 = solve(puzzle, False)
    print(f"Part 1: {sol_part1}")
    print(f"Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
