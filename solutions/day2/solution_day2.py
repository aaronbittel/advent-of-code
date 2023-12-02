with open("puzzle_input_day2", "r") as f:
    input_list = [row.strip() for row in f.readlines()]


values = {"red": 12, "green": 13, "blue": 14}

test_data = [
    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
]


def split_row_in_id_and_values(data):
    game_data = {}
    for row in data:
        game_id = int(row.replace("Game ", "").split(":")[0])
        game_values = row.split(": ")[1].replace(",", "").replace(";", "").split(" ")
        nums = [int(num) for num in game_values[0::2]]
        colors = game_values[1::2]
        tmp_dict = {}
        for color, num in zip(colors, nums):
            if color not in tmp_dict:
                tmp_dict[color] = num
            else:
                if tmp_dict[color] < num:
                    tmp_dict[color] = num
        mult = 1
        for num in tmp_dict.values():
            mult *= num
        game_data[game_id] = mult

    return game_data


total_sum = 0

for game_id, mult in split_row_in_id_and_values(input_list).items():
    total_sum += mult

print(total_sum)
