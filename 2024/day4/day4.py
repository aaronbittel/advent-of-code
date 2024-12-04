import time

DIRECTIONS = [
    (1, 0),
    (-1, 0),
    (0, 1),
    (0, -1),
    (1, 1),
    (-1, -1),
    (1, -1),
    (-1, 1),
]


def check_pattern(matrix, r, c, pattern):
    return sum(
        [
            all(
                matrix.get((r + i * dr, c + i * dc), ".") == s
                for i, s in enumerate(pattern, start=1)
            )
            for dr, dc in DIRECTIONS
        ]
    )


def check_diagonals(matrix, r, c):
    one = matrix.get((r - 1, c - 1), ".") + matrix.get((r + 1, c + 1), ".")
    two = matrix.get((r + 1, c - 1), ".") + matrix.get((r - 1, c + 1), ".")
    return (one == "MS" or one == "SM") and (two == "MS" or two == "SM")


def process_part(matrix, symbol, pattern=None):
    result = 0
    for (r, c), sym in matrix.items():
        if sym != symbol:
            continue
        if pattern:
            result += check_pattern(matrix, r, c, pattern)
        else:
            result += check_diagonals(matrix, r, c)
    return result


def read_input(filename: str, search: str) -> dict:
    with open(filename) as f:
        return {
            (row, col): char
            for row, line in enumerate(f)
            for col, char in enumerate(line.strip())
            if char in search
        }


def main():
    matrix_xmas = read_input("./input.txt", "XMAS")
    matrix_mas = read_input("./input.txt", "MAS")

    start = time.perf_counter()
    p1 = process_part(matrix_xmas, "X", pattern="MAS")
    end = time.perf_counter()
    print(f"Part1: {p1}, took: {end - start:.5f}s")

    start = time.perf_counter()
    p2 = process_part(matrix_mas, "A")
    end = time.perf_counter()
    print(f"Part2: {p2}, took: {end - start:.5f}s")


if __name__ == "__main__":
    main()
