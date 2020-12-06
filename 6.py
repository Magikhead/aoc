from aoc import aoc
from collections import Counter

with open("test/6_input.txt") as f:
    customs_groups = f.read().split("\n\n")
    sum = 0
    for customs_group in customs_groups:
        # strip blank lines
        customs_group = customs_group.strip("\n\n")
        answers = aoc.count_customs_forms(customs_group)
        sum = sum + len(answers)

    print("6-1: %s" % sum)

    sum = 0
    for customs_group in customs_groups:
        # strip blank lines
        customs_group = customs_group.strip("\n\n")
        answers = aoc.count_customs_forms(customs_group)
        # each line represents a form
        num_forms = len(customs_group.split("\n"))
        count = Counter(answers.values())[num_forms]
        sum = sum + count

    print("6-2: %s" % sum)
