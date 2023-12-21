import time
from collections import namedtuple

# convert input into 2-dimensional string array
# convert symbols to two directions, e.g. | (1, 1) connect (0, 0) to (0, 2)
# start at S, walk through both directions at the same time till I find
#   already visited tile


coord = namedtuple("xy", ["x", "y"])


FROM_NORTH = coord(0, 1)
FROM_SOUTH = coord(0, -1)
FROM_WEST = coord(1, 0)
FROM_EAST = coord(-1, 0)


pipes_direction_mapper = {
    "|": {FROM_NORTH: FROM_NORTH, FROM_SOUTH: FROM_SOUTH},
    "-": {FROM_WEST: FROM_WEST, FROM_EAST: FROM_EAST},
    "L": {FROM_NORTH: FROM_WEST, FROM_EAST: FROM_SOUTH},
    "J": {FROM_NORTH: FROM_EAST, FROM_WEST: FROM_SOUTH},
    "7": {FROM_SOUTH: FROM_EAST, FROM_WEST: FROM_NORTH},
    "F": {FROM_SOUTH: FROM_WEST, FROM_EAST: FROM_NORTH},
}


def load(file):
    with open(file, "r") as f:
        return [[sym for sym in row.strip()] for row in f]


def solve_part1(p):
    loop_matrix = [[" " for sym in row] for row in p]
    part1 = 0
    starting_position = coord(0, 0)
    for y, row in enumerate(p):
        for x, next_sym in enumerate(row):
            if next_sym == "S":
                starting_position = coord(x, y)
                break

    starting_directions = [
        direction
        for direction in [coord(0, 1), coord(1, 0), coord(-1, 0), coord(0, -1)]
        if p[xy_add(starting_position, direction).y][
            xy_add(starting_position, direction).x
        ]
        != "."
    ]

    for direction in starting_directions:
        current_position = starting_position
        cur_direction = direction
        loop_matrix[current_position.y][current_position.x] = tuple_get(
            p, current_position
        )
        i = 0
        while True:
            i += 1
            next_pos = xy_add(current_position, cur_direction)
            loop_matrix[next_pos.y][next_pos.x] = tuple_get(p, next_pos)
            next_sym = tuple_get(p, next_pos)
            if next_sym == "S":
                part1 = i // 2
                break
            # print(current_position, cur_direction, next_sym)
            # print(cur_direction, pipes_direction_mapper[next_sym], next_sym)
            if cur_direction in pipes_direction_mapper[next_sym]:
                current_position = next_pos
                cur_direction = pipes_direction_mapper[next_sym][cur_direction]
            else:
                break

    # for row in loop_matrix:
    #    for sym in row:
    #        print(sym, end="")
    #    print()

    return part1, loop_matrix


def solve_part2(loop):
    part2 = 0
    matrix_counter = {}
    for y, row in enumerate(loop_matrix):
        if y in [0, len(loop_matrix) - 1]:
            continue
        for x, sym in enumerate(row):
            counter = 0
            if sym != " ":
                continue
            for x1 in range(x - 1, -1, -1):
                if (x1, y) in matrix_counter:
                    counter += matrix_counter[(x1, y)]
                    break
                if loop[y][x1] in "|JL":
                    counter += 1
            matrix_counter[(x, y)] = counter
            part2 += counter % 2
    return part2


def tuple_get(matrix: list[list[str]], t: coord[int, int]) -> str:
    return matrix[t[1]][t[0]]


def xy_add(coord1: coord, coord2: coord) -> coord[int, int]:
    return coord(coord1.x + coord2.x, coord1.y + coord2.y)


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day10.txt")
    sol_part1, loop_matrix = solve_part1(puzzle)
    sol_part2 = solve_part2(loop_matrix)
    print(f"Part 1: {sol_part1}, Part2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
