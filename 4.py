from aoc import aoc

with open("test/4_input.txt") as f:
    passport_raw = f.read().split("\n\n")
    passport_list = []
    for line in passport_raw:
        line = line.replace("\n", " ")
        line = line.split()

        passport = {}
        for element in line:
            key, value = element.split(":", 1)
            passport[key] = value
        passport_list.append(passport)

    valid_passports = 0
    for passport in passport_list:
        if aoc.is_valid_passport(passport):
            valid_passports = valid_passports + 1

    print("4-1: %s" % valid_passports)

    valid_passports = 0
    for passport in passport_list:
        if aoc.is_valid_passport2(passport):
            valid_passports = valid_passports + 1

    print("4-2: %s" % valid_passports)
