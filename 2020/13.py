from aoc import aoc

with open("test/13_input.txt") as f:
    start_time = int(f.readline())
    bus_schedule_raw = f.read().rstrip().split(",")
    bus_schedule = []
    for bus in bus_schedule_raw:
        if bus.isdigit():
            bus_schedule.append(int(bus))
        else:
            bus_schedule.append(bus)
    next_bus = aoc.find_next_bus(start_time, bus_schedule)
    answer = next_bus[1] * (next_bus[0] - start_time)

    print("13-1: %s" % answer)

    answer = aoc.find_minute_offset_timestamp(bus_schedule)

    print("13-2: %s" % answer)
