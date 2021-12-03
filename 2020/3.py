from aoc import aoc

with open("test/3_input.txt") as f:
    toboggan_map_raw = f.read().splitlines()
    toboggan_map = list(map(list, toboggan_map_raw))

    tree_count = aoc.toboggan_map_count_trees(toboggan_map, (3, 1))
    print("3-1: %s" % tree_count)

    slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
    total = 1
    for slope in slopes:
        total = total * aoc.toboggan_map_count_trees(toboggan_map, slope)
    print("3-2: %s" % total)
