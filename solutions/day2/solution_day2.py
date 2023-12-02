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
        game_data[game_id] = []
        games = row.split(": ")[1].split("; ")
        for game in games:
            tmp_dict = {}
            for num in game.split(", "):
                num_color_pair = num.split(" ")
                tmp_dict[num_color_pair[1]] = int(num_color_pair[0])
            game_data[game_id].append(tmp_dict)
    return game_data


total_sum = 0

for game_id, game_values in split_row_in_id_and_values(input_list).items():
    to_add = True
    for game in game_values:
        for color, num in values.items():
            if values[color] < game.get(color, -1):
                to_add = False
    total_sum += game_id if to_add else 0


print(total_sum)
