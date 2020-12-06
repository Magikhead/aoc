from aoc import aoc

with open("test/5_input.txt") as f:
    seat_codes = f.read().splitlines()
    seat_ids = []
    for seat_code in seat_codes:
        seat = aoc.decode_seat(seat_code)
        seat_ids.append(aoc.seat_id(seat))

    print("5-1: %s" % max(seat_ids))

    for x in range(min(seat_ids), max(seat_ids)):
        if x not in seat_ids:
            my_seat = x

    print("5-2: %s" % my_seat)
