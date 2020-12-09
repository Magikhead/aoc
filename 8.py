from aoc import aoc

with open("test/8_input.txt") as f:
    instructions_raw = f.read().splitlines()

    program = []
    for line in instructions_raw:
        instruction = aoc.parse_instruction(line)
        program.append(instruction)

    accumulator, ptr_history = aoc.run_program(program)

    print("8-1: %s" % accumulator)

    suspect = 0
    while True:
        program[suspect] = aoc.flip_instruction(program[suspect])
        accumulator, ptr_history = aoc.run_program(program)
        if ptr_history[-1] == len(program) - 1:
            # program terminates normally
            break
        else:
            program[suspect] = aoc.flip_instruction(program[suspect])
            suspect = suspect + 1

    print("8-2: %s" % accumulator)
