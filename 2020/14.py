from aoc import aoc

with open("test/14_input.txt") as f:
    instructions = f.read().splitlines()

    emulator = aoc.Emulator()
    for instruction in instructions:
        emulator.execute(instruction)
    sum = 0
    for addr, memory in emulator.memory.items():
        sum = sum + memory

    print("14-1: %s" % sum)

    emulator = aoc.Emulator2()
    for instruction in instructions:
        emulator.execute(instruction)
    sum = 0
    for addr, memory in emulator.memory.items():
        sum = sum + memory

    print("14-2: %s" % sum)
