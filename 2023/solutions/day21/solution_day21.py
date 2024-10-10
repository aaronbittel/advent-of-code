import time


UP = (0, -1)
UP_RIGHT = (1, -1)
RIGHT = (1, 0)
DOWN_RIGHT = (1, 1)
DOWN = (0, 1)
DOWN_LEFT = (-1, 1)
LEFT = (-1, 0)
UP_LEFT = (-1, -1)


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p):
    start = next(
        (x, y) for y, row in enumerate(p) for x, sym in enumerate(row) if sym == "S"
    )

    steps = 64

    grid = {
        (x, y): c
        for y, row in enumerate(p)
        for x, c in enumerate(row)
        if distance((x, y), start) <= steps + 1
    }

    garden(grid, start, steps)

    print_grid(grid)

    return sum(1 for c in grid.values() if c.isdigit() and int(c) % 2 == 0) + 1


def print_grid(grid):
    max_x, max_y = max(x for x, _ in grid), max(y for _, y in grid)

    for y in range(max_y):
        for x in range(max_x):
            c = grid.get((x, y), None)
            if c is None:
                print("  ", end=" ")
            elif len(grid[(x, y)]) == 1:
                print(" " + grid[(x, y)], end=" ")
            else:
                print(grid[(x, y)], end=" ")
        print()
    print()


def distance(p1: tuple[int, int], p2: tuple[int, int]):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def garden(grid, start, steps=64):
    index = 1
    q = [(start, (0, 0), steps)]
    while q:
        index += 1
        pos, (dx, dy), steps_left = q.pop(0)
        for dir2 in {RIGHT, LEFT, UP, DOWN} - {(dx, dy), (-dx, -dy)}:
            steps_left2 = steps_left
            (x2, y2), (dx2, dy2) = pos, dir2
            while True:
                next_pos = (x2 + dx2, y2 + dy2)
                next_sym = grid.get(next_pos, None)
                if (
                    next_sym is None
                    or next_sym == "#"
                    or steps_left2 <= 0
                    or (next_sym.isdigit() and int(next_sym) >= steps_left2 - 1)
                ):
                    break
                steps_left2 -= 1
                grid[*next_pos] = str(steps_left2)
                if steps_left2 > 0:
                    q.append((next_pos, dir2, steps_left2))
                x2 += dx2
                y2 += dy2

        # print_grid(grid)


if __name__ == "__main__":
    time_start = time.perf_counter()
    solution = solve(load("puzzle_input_day21.txt"))
    print(f"Part 1: {solution}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
