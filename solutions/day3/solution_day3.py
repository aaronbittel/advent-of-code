with open("puzzle_input_day3", "r") as f:
    input_list = [line.strip() for line in f.readlines()]


test_case = [
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


DIRECTIONS = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)]


def check_row_col(input_list, row, col):
    max_col_length = len(input_list[0]) - 1
    max_row_length = len(input_list) - 1
    return (0 <= row <= max_row_length) and (0 <= col <= max_col_length)


def get_adjacent_number_positions(input_list, row_index, col_index):
    adjacent_number_position = []
    for direction in DIRECTIONS:
        d_row, d_col = direction
        row_to_look = row_index + d_row
        col_to_look = col_index + d_col
        if not check_row_col(input_list, row_to_look, col_to_look):
            continue
        symbol = input_list[row_to_look][col_to_look]
        if symbol.isdigit():
            adjacent_number_position.append((row_to_look, col_to_look))
    return adjacent_number_position


def get_number(input_list: list[str], number_positions):
    already_added_numbers = []
    for number_pos in number_positions:
        row, col = number_pos[0], number_pos[1]
        start, end = col, col
        while start - 1 >= 0:
            if input_list[row][start - 1].isdigit():
                start -= 1
            else:
                break
        while end + 1 < len(input_list[0]):
            if input_list[row][end + 1].isdigit():
                end += 1
            else:
                break
        number = int(input_list[row][start : end + 1])
        if number in already_added_numbers:
            continue
        already_added_numbers.append(number)
    return already_added_numbers


def main(input_list):
    total_sum = 0
    for row_index, row in enumerate(input_list):
        for col_index, symbol in enumerate(row):
            if symbol.isdigit() or symbol == ".":
                continue
            number_positions = get_adjacent_number_positions(
                input_list, row_index, col_index
            )
            if not number_positions:
                continue
            numbers = get_number(input_list, number_positions)
            for num in numbers:
                total_sum += num
    print(total_sum)


if __name__ == "__main__":
    main(input_list)
