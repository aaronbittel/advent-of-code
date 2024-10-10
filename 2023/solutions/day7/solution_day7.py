import time
from collections import Counter, namedtuple
from enum import Enum


camel_card = namedtuple("camel_card", ["hand", "bid", "best_hand"])


class HandRank(Enum):
    FIVE_OF_A_KIND = 0
    FOUR_OF_A_KIND = 1
    FULL_HOUSE = 2
    THREES_OF_A_KIND = 3
    TWO_PAIR = 4
    ONE_PAIR = 5
    HIGH_CARD = 6


# A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
card_strength_part1 = {
    "A": 13,
    "K": 12,
    "Q": 11,
    "J": 10,
    "T": 9,
    "9": 8,
    "8": 7,
    "7": 6,
    "6": 5,
    "5": 4,
    "4": 3,
    "3": 2,
    "2": 1,
}

card_strength_part2 = {
    "A": 13,
    "K": 12,
    "Q": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2,
    "J": 1,
}


def load(file):
    with open(file, "r") as f:
        return [row.strip() for row in f]


def solve(p, part1: bool = True):
    solution = 0
    hands = [row.split(" ")[0] for row in p]
    bids = list(map(int, [row.split(" ")[1] for row in p]))
    camel_cards = [
        camel_card(hand, bid, calc_best_hand(hand)) for hand, bid in zip(hands, bids)
    ]
    camel_cards_type_mapping = [[] for _ in range(7)]
    for cml_card in camel_cards:
        if part1:
            camel_cards_type_mapping[get_type(cml_card.hand)].append(cml_card)
        else:
            camel_cards_type_mapping[get_type(cml_card.best_hand)].append(cml_card)
    for type_camel_cards in camel_cards_type_mapping:
        if part1:
            type_camel_cards.sort(key=sort_same_group_part1, reverse=True)
        else:
            type_camel_cards.sort(key=sort_same_group_part2, reverse=True)
    amount_of_hands = len(hands)
    for type_camel_cards in camel_cards_type_mapping:
        for cml_card in type_camel_cards:
            solution += amount_of_hands * cml_card.bid
            amount_of_hands -= 1
    return solution


def calc_best_hand(hand: str):
    occurrences = Counter(hand)
    if "J" not in occurrences:
        return hand
    if occurrences["J"] == 5:
        return "AAAAA"
    card_occ = namedtuple("card_occ", ["card", "occ"])
    jokers_replace_card = card_occ("J", -1)
    for card, occ in occurrences.items():
        if card == "J":
            continue
        if occ > jokers_replace_card.occ:
            jokers_replace_card = card_occ(card, occ)
        elif (
            occ == card_occ.occ
            and card_strength_part2[card]
            > card_strength_part2[jokers_replace_card.card]
        ):
            jokers_replace_card = card_occ(card, occ)
    return hand.replace("J", jokers_replace_card.card)


def get_type(hand):
    occurrences = Counter(hand)
    match len(Counter(hand)):
        case 1:
            return HandRank.FIVE_OF_A_KIND.value
        case 2:
            return (
                HandRank.FULL_HOUSE.value
                if all(o >= 2 for o in occurrences.values())
                else HandRank.FOUR_OF_A_KIND.value
            )
        case 3:
            return (
                HandRank.THREES_OF_A_KIND.value
                if any(o == 3 for o in occurrences.values())
                else HandRank.TWO_PAIR.value
            )
        case 4:
            return HandRank.ONE_PAIR.value
        case 5:
            return HandRank.HIGH_CARD.value
        case _:
            return float("inf")


def sort_same_group_part1(type_camel_cards: camel_card):
    return [card_strength_part1[card] for card in type_camel_cards.hand]


def sort_same_group_part2(type_camel_cards: camel_card):
    return [card_strength_part2[card] for card in type_camel_cards.hand]


if __name__ == "__main__":
    time_start = time.perf_counter()
    puzzle = load("puzzle_input_day7.txt")
    solution_1 = solve(puzzle)
    solution_2 = solve(puzzle, part1=False)
    print(f"Part1: {solution_1}, Part2: {solution_2}")
    print(f"Solved in {time.perf_counter() - time_start:.5f} Sec.")
