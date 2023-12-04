with open("puzzle_input_day4", "r") as f:
    input_file = [line.strip() for line in f.readlines()]


test_case_0 = [
    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
    "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
    "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
    "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
    "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
    "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
]


def convert_str_to_int(lst):
    new_list = []
    for l in lst:
        new_list.append([int(item) for item in l if len(item) > 0])
    return new_list


def check_if_winning(winning_numbers, my_numbers):
    winning_count = []
    for index, lst in enumerate(my_numbers):
        count = 0
        for num in lst:
            if num == "":
                continue
            if num in winning_numbers[index]:
                count += 1
        winning_count.append(count)
    return winning_count


def solution(input_list):
    winning_numbers_list = []
    my_numbers_list = []
    for line in input_list:
        numbers = line.split(": ")[1]
        winning_numbers_string = numbers.split(" | ")[0]
        my_numbers_string = numbers.split(" | ")[1]
        # print(winning_numbers_string, my_numbers_string, sep="\t\t")

        winning_numbers_list.append(winning_numbers_string.split(" "))
        my_numbers_list.append(my_numbers_string.split(" "))

    winning_numbers_list = convert_str_to_int(winning_numbers_list)
    my_numbers_list = convert_str_to_int(my_numbers_list)

    winning_count = check_if_winning(winning_numbers_list, my_numbers_list)
    total_sum = 0
    print(winning_count)
    for index, count in enumerate(winning_count):
        total_sum += 2 ** (count - 1) if count != 0 else 0
    print(total_sum)


def solution_2(input_list):
    scratchcard_copies = {i: 1 for i in range(1, len(input_list) + 1)}
    print(scratchcard_copies)

    winning_numbers_list = []
    my_numbers_list = []
    for line in input_list:
        numbers = line.split(": ")[1]
        winning_numbers_string = numbers.split(" | ")[0]
        my_numbers_string = numbers.split(" | ")[1]
        # print(winning_numbers_string, my_numbers_string, sep="\t\t")

        winning_numbers_list.append(winning_numbers_string.split(" "))
        my_numbers_list.append(my_numbers_string.split(" "))

    winning_count = check_if_winning(winning_numbers_list, my_numbers_list)
    card_winning_count = {i: wins for i, wins in enumerate(winning_count, start=1)}
    print(card_winning_count)

    max_length = len(scratchcard_copies)

    for card_no, card_wins in zip(scratchcard_copies, card_winning_count.values()):
        print(card_no)
        for _ in range(scratchcard_copies[card_no]):
            for i in range(1, card_wins + 1):
                if card_no + i > max_length:
                    break
                scratchcard_copies[card_no + i] += 1
            # print(" -> ", scratchcard_copies)

    print(sum(scratchcard_copies.values()))


if __name__ == "__main__":
    # solution(input_file)
    solution_2(input_file)
