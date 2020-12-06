import itertools as it
import numpy
import re


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


def is_valid_passport(passport):
    required_keys = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
    for key in required_keys:
        if key not in passport:
            return False
    return True


def is_valid_num(num, **kwargs):
    length = kwargs.get("length", None)
    min = kwargs.get("min", None)
    max = kwargs.get("max", None)

    if length is not None:
        if len(str(num)) != length:
            return False
    if min is not None:
        if num < min:
            return False
    if max is not None:
        if num > max:
            return False
    return True


def is_valid_birth_year(year):
    return is_valid_num(year, length=4, min=1920, max=2002)


def is_valid_issue_year(year):
    return is_valid_num(year, length=4, min=2010, max=2020)


def is_valid_expiration_year(year):
    return is_valid_num(year, length=4, min=2020, max=2030)


def is_valid_height(height):
    pattern = "^([0-9]+)(cm|in)$"
    result = re.search(pattern, height)
    if result is None:
        return False
    if result.group(2) == "cm":
        if not is_valid_num(int(result.group(1)), min=150, max=193):
            return False
    elif result.group(2) == "in":
        if not is_valid_num(int(result.group(1)), min=59, max=76):
            return False
    else:
        return False

    return True


def is_valid_hair_color(hair_color):
    pattern = "^#[0-9a-f]{6}$"
    result = re.search(pattern, hair_color)
    if result is None:
        return False
    return True


def is_valid_eye_color(eye_color):
    eye_colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
    return eye_color in eye_colors


def is_valid_passport_id(passport_id):
    pattern = "^[0-9]{9}$"
    result = re.search(pattern, passport_id)
    if result is None:
        return False
    return True


def is_valid_passport2(passport):
    if not is_valid_passport(passport):
        return False

    if not is_valid_birth_year(int(passport["byr"])):
        return False
    if not is_valid_issue_year(int(passport["iyr"])):
        return False
    if not is_valid_expiration_year(int(passport["eyr"])):
        return False
    if not is_valid_height(passport["hgt"]):
        return False
    if not is_valid_hair_color(passport["hcl"]):
        return False
    if not is_valid_eye_color(passport["ecl"]):
        return False
    if not is_valid_passport_id(passport["pid"]):
        return False

    return True


def decode_seat(seat_code):
    seat_code = seat_code.replace("F", "0")
    seat_code = seat_code.replace("B", "1")
    seat_code = seat_code.replace("L", "0")
    seat_code = seat_code.replace("R", "1")
    row = seat_code[0:7]
    column = seat_code[-3:]
    # convert from binary to int
    row = int(row, 2)
    column = int(column, 2)
    return (row, column)


def seat_id(seat):
    return seat[0] * 8 + seat[1]


def count_customs_forms(forms):
    answers = {}
    for form in forms:
        for question in form:
            if question.isalpha():
                answers[question] = answers.get(question, 0) + 1
    return answers
