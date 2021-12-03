from aoc import aoc

with open("test/1_input.txt") as f:
    expense_report = f.read().splitlines()
    expense_report = [int(i) for i in expense_report]

    result = aoc.fix_expense_report(expense_report)
    print("1-1: %s" % result)

    result = aoc.fix_expense_report(expense_report, 3)
    print("1-2: %s" % result)
