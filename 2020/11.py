from aoc import aoc

with open("test/11_input.txt") as f:
    seats_raw = f.read().splitlines()
    seats = []
    for row in seats_raw:
        row = list(row)
        seats.append(row)

    waiting_area = aoc.WaitingArea(seats)
    steps = waiting_area.run()

    print("11-1: %s" % waiting_area.count_occupied())

    waiting_area = aoc.WaitingArea(seats, max_depth=None, neighbor_threshold=5)
    steps = waiting_area.run()

    print("11-2: %s" % waiting_area.count_occupied())
