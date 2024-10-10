with open("puzzle_input_day3.txt", "r") as f:
    input_file = [line.strip() for line in f.readlines()]

test_case_0 = [
    "467..114..",
    "...*......",
    "..35..633.",
    "......#...",
    "617*......",
    ".....+.58.",
    "..592.....",
    "......755.",
    "...$.*....",
    ".664.598..",
]

test_case_1 = [
    "467.114..1",
    "...*......",
    "..35..633.",
    "......#...",
    "617*......",
    ".....+.58.",
    "..592.....",
    "......755.",
    "...$&*....",
    ".664.598..",
]

test_case_2 = [
    "......755.",
    "..6$.*12..",
    ".6.6.4.59.",
]


DIRECTIONS = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)]


def check_row_col(input_list, row, col):
    max_col_length = len(input_list[0]) - 1
    max_row_length = len(input_list) - 1
    return (0 <= row <= max_row_length) and (0 <= col <= max_col_length)


def solution(input_list):
    numbers = []
    for row_index, row in enumerate(input_list):
        for col_index, symbol in enumerate(row):
            tmp_numbers = []
            if symbol != "*":
                continue
            visited_cords = []
            for direction in DIRECTIONS:
                row_to_look, col_to_look = (
                    row_index + direction[0],
                    col_index + direction[1],
                )
                if not check_row_col(input_list, row_to_look, col_to_look):
                    continue
                if (row_to_look, col_to_look) in visited_cords or not input_list[
                    row_to_look
                ][col_to_look].isdigit():
                    continue
                start, end = col_to_look, col_to_look
                while start - 1 >= 0:
                    if not input_list[row_to_look][start - 1].isdigit():
                        visited_cords.append((row_to_look, start))
                        break
                    start -= 1
                    visited_cords.append((row_to_look, start))
                while end + 1 < len(row):
                    if not input_list[row_to_look][end + 1].isdigit():
                        visited_cords.append((row_to_look, end))
                        break
                    end += 1
                    visited_cords.append((row_to_look, end))
                number = int(input_list[row_to_look][start : end + 1])
                tmp_numbers.append(number)
                print(f"{symbol} -> {number}")
            if len(tmp_numbers) == 2:
                numbers.append(tmp_numbers[0] * tmp_numbers[1])
    return numbers


if __name__ == "__main__":
    total_sum = 0
    numbers = solution(input_file)
    for num in numbers:
        total_sum += num
    print(total_sum)
