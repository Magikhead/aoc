from aoc import aoc

with open("test/7_input.txt") as f:
    bag_rules_raw = f.read().splitlines()
    bag_rules = {}
    for line in bag_rules_raw:
        bag_rule = aoc.parse_bag_line(line)
        bag_rules[bag_rule[0]] = bag_rule[1]

    colors_with_gold_bags = []
    for bag_color in bag_rules.keys():
        if aoc.check_bag_color("shiny gold", bag_color, bag_rules):
            colors_with_gold_bags.append(bag_color)
    # shiny gold cannot be the outter-most
    total_colors = len(colors_with_gold_bags) - 1
    print("7-1: %s" % total_colors)

    # do not count the outter-most bag in the total
    num_bags = aoc.count_bags("shiny gold", bag_rules) - 1
    print("7-2: %s" % num_bags)
