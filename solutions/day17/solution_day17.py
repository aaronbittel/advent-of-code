import time
import heapq

RIGHT = (1, 0)
LEFT = (-1, 0)
UP = (0, -1)
DOWN = (0, 1)


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    grid = {(x, y): int(c) for y, row in enumerate(p) for x, c in enumerate(row)}
    part1 = path(max(grid), grid, 1, 3)
    part2 = path(max(grid), grid, 4, 10)
    return part1, part2


def path(target, grid, least, most):
    q, visited = [(0, (0, 0), (0, 0))], set()

    while q:
        heat, pos, direction = heapq.heappop(q)
        if pos == target:
            return heat
        if (pos, direction) in visited:
            continue
        visited.add((pos, direction))

        (x, y), (dx, dy) = pos, direction
        for dir2 in {RIGHT, LEFT, UP, DOWN} - {
            (dx, dy),
            (-dx, -dy),
        }:
            dx2, dy2 = dir2
            h = heat
            for mul in range(1, most + 1):
                pos2 = (x + dx2 * mul, y + dy2 * mul)
                if pos2 not in grid:
                    break
                h += grid[pos2]
                if mul < least:
                    continue
                heapq.heappush(q, (h, pos2, dir2))


if __name__ == "__main__":
    time_start = time.perf_counter()
    sol_part1, sol_part2 = solve(load("puzzle_input_day17.txt"))
    print(f"Part 1: {sol_part1}, Part 2: {sol_part2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
