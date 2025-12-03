from pathlib import Path
import sys


def repeated_part1(s: str) -> bool:
    if len(s) % 2 == 1:
        return False
    mid = len(s) // 2
    return s[:mid] == s[mid:]


def repeated_part2(s: str) -> bool:
    for i in range(1, (len(s) // 2) + 1):
        pattern = s[:i]
        l_pattern = len(pattern)
        if len(s) % l_pattern != 0:
            continue
        for j in range(1, len(s) // l_pattern):
            test = s[j * l_pattern : (j + 1) * l_pattern]
            if pattern != test:
                break
        else:
            return True
    return False


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print(f"USAGE: {sys.argv[0]} <input-file>", file=sys.stderr)
        sys.exit(1)

    content = Path(sys.argv[1]).read_text()
    pairs = content.split(",")
    sol_part1 = 0
    for pair in pairs:
        start, stop = pair.split("-")
        for n in range(int(start), int(stop) + 1):
            if repeated_part1(str(n)):
                sol_part1 += n

    sol_part2 = 0
    for pair in pairs:
        start, stop = pair.split("-")
        for n in range(int(start), int(stop) + 1):
            if repeated_part2(str(n)):
                sol_part2 += n

    print(f"Sol Part1: {sol_part1}")
    print(f"Sol Part2: {sol_part2}")
