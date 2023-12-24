import time
from dataclasses import dataclass
from typing import Set


@dataclass
class Point:
    x: int
    y: int

    def __add__(self, other):
        return Point(self.x + other.x, self.y + other.y)

    def __hash__(self):
        return hash(self.x) + hash(self.y)


RIGHT = Point(1, 0)
LEFT = Point(-1, 0)
UP = Point(0, -1)
DOWN = Point(0, 1)


def load(file):
    with open(file, "r") as f:
        return [[c for c in row.strip()] for row in f]


def solve(p, part1):
    if part1:
        visited = set()
        beam(p, visited, Point(-1, 0), RIGHT)
        return len(set(pos for pos, _ in visited))
    else:
        part2 = 0

        for i in range(len(p[0])):
            visited = set()
            beam(p, visited, Point(i, -1), DOWN)
            part2 = max(part2, (len(set(pos for pos, _ in visited))))
            visited = set()
            beam(p, visited, Point(i, len(p)), UP)
            part2 = max(part2, (len(set(pos for pos, _ in visited))))

            print(f"#1: {(i + 1) / len(p[0]) * 100:.2f} %")

        for i in range(len(p)):
            visited = set()
            beam(p, visited, Point(-1, i), RIGHT)
            part2 = max(part2, (len(set(pos for pos, _ in visited))))
            visited = set()
            beam(p, visited, Point(len(p[0]), i), LEFT)
            part2 = max(part2, (len(set(pos for pos, _ in visited))))

            print(f"#2: {(i + 1) / len(p) * 100:.2f} %")

        return part2


def beam(
    grid: list[list[str]],
    visited: Set[tuple[Point, Point]],
    pos: Point,
    direction: Point,
):
    next_pos: Point = pos + direction

    while (
        0 <= next_pos.x < len(grid[0])
        and 0 <= next_pos.y < len(grid)
        and tuple((next_pos, direction)) not in visited
    ):
        visited.add((next_pos, direction))
        match grid[next_pos.y][next_pos.x]:
            case ".":
                next_pos += direction
            case "|":
                if direction in (RIGHT, LEFT):
                    beam(grid, visited, next_pos, UP)
                    beam(grid, visited, next_pos, DOWN)
                else:
                    next_pos += direction
            case "-":
                if direction in (RIGHT, LEFT):
                    next_pos += direction
                else:
                    beam(grid, visited, next_pos, RIGHT)
                    beam(grid, visited, next_pos, LEFT)
            case "/":
                direction = Point(-direction.y, -direction.x)
                next_pos += direction
            case "\\":
                direction = Point(direction.y, direction.x)
                next_pos += direction


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day16.txt")
    sol_part1 = solve(puzzle, True)
    sol_part2 = solve(puzzle, False)
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
