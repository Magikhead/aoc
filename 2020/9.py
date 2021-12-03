from aoc import aoc

with open("test/9_input.txt") as f:
    ciphertext_raw = f.read().splitlines()
    ciphertext = [int(x) for x in ciphertext_raw]
    number = aoc.xmas_find_non_sum(ciphertext, 25)

    print("9-1: %s" % number)

    weakness = aoc.xmas_find_weakness(number, ciphertext)
    print("9-2: %s" % weakness)
