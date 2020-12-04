from aoc import aoc

with open("test/2_input.txt") as f:
    password_list = []
    password_list_raw = f.read().splitlines()
    for entry in password_list_raw:
        entry = entry.replace(":", "")
        entry = entry.replace("-", " ")
        entry = entry.split()
        password_list.append(entry)

    count = 0
    for entry in password_list:
        if aoc.check_password(*entry):
            count = count + 1
    print("2-1: %s" % count)

    count = 0
    for entry in password_list:
        if aoc.check_toboggan_corporate_password(entry[0:2], *entry[2:]):
            count = count + 1
    print("2-2: %s" % count)
