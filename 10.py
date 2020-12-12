from aoc import aoc

with open("test/10_input.txt") as f:
    adapters_list_raw = f.read().splitlines()
    adapters_list = [int(x) for x in adapters_list_raw]

    # sort the adapters and add charging outlet and device joltages
    adapters_list.sort()
    device_joltage = adapters_list[-1] + 3
    adapters_list.append(device_joltage)
    adapters_list.insert(0, 0)

    joltage_array = aoc.count_jolt_differences(adapters_list)
    answer = joltage_array[1] * joltage_array[3]

    print("10-1: %s" % answer)

    adapters_list.sort()
    answer = aoc.count_adapter_combinations(adapters_list)

    print("10-2: %s" % answer)
