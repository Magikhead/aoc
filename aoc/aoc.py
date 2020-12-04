import itertools as it
import numpy


def fix_expense_report(expenses, num_entries=2):
    for i in list(it.permutations(expenses, num_entries)):
        if sum(i) == 2020:
            return numpy.prod(i)
    return 0


def check_password(min, max, letter, password):
    min = int(min)
    max = int(max)
    occurrences = password.count(letter)
    return occurrences >= min and occurrences <= max


def check_toboggan_corporate_password(indexes, letter, password):
    count = 0
    for index in indexes:
        index = int(index)
        if password[index - 1] == letter:
            count = count + 1
    return count == 1


def toboggan_map_count_trees(map, slope):
    tree_count = 0
    tree = "#"
    width = len(map[0])

    # start in top left corner of map (0,0)
    x = 0
    y = 0

    while y < len(map):
        if map[y][x % width] == tree:
            tree_count = tree_count + 1
        x = (x + slope[0]) % width  # map repeats indefinitely to the right
        y = y + slope[1]
    return tree_count
