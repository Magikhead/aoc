from aoc import aoc

with open("test/12_input.txt") as f:
    navigation_instructions = f.read().splitlines()

    ferry = aoc.Ferry()
    for instruction in navigation_instructions:
        ferry.move(instruction)

    coordinates = ferry.get_coordinates()
    manhattan_distance = abs(coordinates[0]) + abs(coordinates[1])

    print("12-1: %s" % manhattan_distance)

    ferry = aoc.WaypointFerry()
    for instruction in navigation_instructions:
        ferry.move(instruction)

    coordinates = ferry.get_coordinates()
    manhattan_distance = abs(coordinates[0]) + abs(coordinates[1])

    print("12-1: %s" % manhattan_distance)
