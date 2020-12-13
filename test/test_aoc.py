from aoc import aoc
import unittest
import pdb


class TestAoCDay1(unittest.TestCase):
    def test_fix_expense_report(self):
        self.assertEqual(
            514579, aoc.fix_expense_report([1721, 979, 366, 299, 675, 1456])
        )

    def test_fix_expense_report(self):
        self.assertEqual(
            241861950, aoc.fix_expense_report([1721, 979, 366, 299, 675, 1456], 3)
        )


class TestAoCDay2(unittest.TestCase):
    def test_check_password(self):
        self.assertEqual(True, aoc.check_password(1, 3, "a", "abcde"))
        self.assertEqual(False, aoc.check_password(1, 3, "b", "cdefg"))
        self.assertEqual(True, aoc.check_password(2, 9, "c", "ccccccccc"))

    def test_check_toboggan_corporate_password(self):
        self.assertEqual(
            True, aoc.check_toboggan_corporate_password([1, 3], "a", "abcde")
        )
        self.assertEqual(
            False, aoc.check_toboggan_corporate_password([1, 3], "b", "cdefg")
        )
        self.assertEqual(
            False, aoc.check_toboggan_corporate_password([2, 9], "c", "ccccccccc")
        )


class TestAoCDay3(unittest.TestCase):
    def test_toboggan_map_count_trees(self):
        map_1 = [[".", "."], [".", "#"]]
        self.assertEqual(1, aoc.toboggan_map_count_trees(map_1, (1, 1)))

        map_2 = [
            [".", ".", ".", "."],
            [".", "#", ".", "#"],
            [".", ".", "#", "."],
            [".", "#", ".", "#"],
        ]
        self.assertEqual(3, aoc.toboggan_map_count_trees(map_2, (1, 1)))


class TestAoCDay4(unittest.TestCase):
    def test_valid_passport(self):
        passport = {
            "ecl": "gry",
            "pid": "860033327",
            "eyr": "2020",
            "hcl": "#fffffd",
            "byr": "1937",
            "iyr": "2017",
            "cid": "147",
            "hgt": "183cm",
        }
        self.assertEqual(True, aoc.is_valid_passport(passport))

    def test_invalid_passport(self):
        passport = {
            "iyr": "2013",
            "ecl": "amb",
            "cid": "350",
            "eyr": "2023",
            "pid": "028048884",
            "hcl": "#cfa07d",
            "byr": "1929",
        }
        self.assertEqual(False, aoc.is_valid_passport(passport))

    def test_is_valid_birth_year(self):
        self.assertEqual(True, aoc.is_valid_birth_year(2002))
        self.assertEqual(False, aoc.is_valid_birth_year(2003))

    def test_is_valid_issue_year(self):
        self.assertEqual(True, aoc.is_valid_issue_year(2010))
        self.assertEqual(False, aoc.is_valid_issue_year(1990))

    def test_is_valid_expiration_year(self):
        self.assertEqual(True, aoc.is_valid_expiration_year(2020))
        self.assertEqual(False, aoc.is_valid_expiration_year(1990))

    def test_is_valid_height(self):
        self.assertEqual(True, aoc.is_valid_height("60in"))
        self.assertEqual(True, aoc.is_valid_height("190cm"))
        self.assertEqual(False, aoc.is_valid_height("190in"))
        self.assertEqual(False, aoc.is_valid_height("90"))

    def test_is_valid_hair_color(self):
        self.assertEqual(True, aoc.is_valid_hair_color("#123abc"))
        self.assertEqual(False, aoc.is_valid_hair_color("#123abz"))
        self.assertEqual(False, aoc.is_valid_hair_color("123abc"))

    def test_is_valid_eye_color(self):
        self.assertEqual(True, aoc.is_valid_eye_color("brn"))
        self.assertEqual(False, aoc.is_valid_eye_color("wat"))

    def test_is_valid_passport_id(self):
        self.assertEqual(True, aoc.is_valid_passport_id("000000001"))
        self.assertEqual(False, aoc.is_valid_passport_id("0123456789"))


class TestAoCDay5(unittest.TestCase):
    def test_decode_seat(self):
        self.assertEqual((44, 5), aoc.decode_seat("FBFBBFFRLR"))

    def test_seat_id(self):
        self.assertEqual(357, aoc.seat_id((44, 5)))


class TestAoCDay6(unittest.TestCase):
    def test_count_custom_forms(self):
        self.assertEqual(3, len(aoc.count_customs_forms(["abc"])))
        self.assertEqual(3, len(aoc.count_customs_forms(["a", "b", "c"])))
        self.assertEqual(1, len(aoc.count_customs_forms(["a", "a", "a"])))
        self.assertEqual(1, len(aoc.count_customs_forms(["b"])))


class TestAoCDay7(unittest.TestCase):
    def test_parse_bag_line(self):
        self.assertEqual(
            ("vibrant bronze", ["2 dim fuchsia"]),
            aoc.parse_bag_line("vibrant bronze bags contain 2 dim fuchsia bags."),
        )
        self.assertEqual(
            ("light red", ["1 bright white", "2 muted yellow"]),
            aoc.parse_bag_line(
                "light red bags contain 1 bright white bag, 2 muted yellow bags."
            ),
        )
        self.assertEqual(
            ("dotted black", ["no other"]),
            aoc.parse_bag_line("dotted black bags contain no other bags."),
        )

    def test_check_bag_color(self):
        bag_rules = {"dark chartreuse": ["1 dark maroon", "5 shiny gold"]}
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "shiny gold", bag_rules)
        )
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "dark chartreuse", bag_rules)
        )

    def test_check_bag_color_many_rules(self):
        bag_rules = {
            "light red": ["1 bright white", "2 muted yellow"],
            "dark orange": ["3 bright white", "4 muted yellow"],
            "bright white": ["1 shiny gold"],
            "muted yellow": ["2 shiny gold", "9 faded blue"],
            "shiny gold": ["1 dark olive", "2 vibrant plum"],
            "dark olive": ["3 faded blue", "4 dotted black"],
            "vibrant plum": ["5 faded blue", "6 dotted black"],
            "faded blue": ["no other"],
            "dotted black": ["no other"],
        }
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "bright white", bag_rules)
        )
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "muted yellow", bag_rules)
        )
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "dark orange", bag_rules)
        )
        self.assertEqual(
            True, aoc.check_bag_color("shiny gold", "light red", bag_rules)
        )
        self.assertEqual(
            False, aoc.check_bag_color("shiny gold", "dark olive", bag_rules)
        )
        self.assertEqual(
            False, aoc.check_bag_color("shiny gold", "dotted black", bag_rules)
        )


class TestAoCDay8(unittest.TestCase):
    def test_run_program(self):
        program = [aoc.Instruction(aoc.OpCode.nop, 0)]
        self.assertEqual((0, [0]), aoc.run_program(program))

    def test_parse_instruction(self):
        instruction = aoc.Instruction(aoc.OpCode.nop, 0)
        self.assertEqual(instruction, aoc.parse_instruction("nop +0"))

    def test_flip_instruction(self):
        instruction = aoc.Instruction(aoc.OpCode.nop, 0)
        instruction = aoc.flip_instruction(instruction)
        self.assertEqual(instruction, aoc.Instruction(aoc.OpCode.jmp, 0))


class TestAoCDay9(unittest.TestCase):
    def test_find_sum_pair(self):
        self.assertEqual((15, 25), aoc.find_sum_pair(40, [35, 20, 15, 25]))

    def test_xmas_find_non_sum(self):
        ciphertext = [
            35,
            20,
            15,
            25,
            47,
            40,
            62,
            55,
            65,
            95,
            102,
            117,
            150,
            182,
            127,
            219,
            299,
            277,
            309,
            576,
        ]
        self.assertEqual(127, aoc.xmas_find_non_sum(ciphertext, 5))

    def test_xmas_find_weakness(self):
        ciphertext = [
            35,
            20,
            15,
            25,
            47,
            40,
            62,
            55,
            65,
            95,
            102,
            117,
            150,
            182,
            127,
            219,
            299,
            277,
            309,
            576,
        ]
        self.assertEqual(62, aoc.xmas_find_weakness(127, ciphertext))


class TestAoCDay10(unittest.TestCase):
    def test_count_jolt_differences(self):
        adapters = [0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22]
        self.assertDictEqual({1: 7, 3: 5}, aoc.count_jolt_differences(adapters))

    def test_find_next_adapter(self):
        adapters = [0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22]
        self.assertEqual([1], aoc.find_next_adapter(0, adapters, 3))

    def test_count_adapter_combinations(self):
        adapters = [0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22]
        self.assertEqual(8, aoc.count_adapter_combinations(adapters))


class TestAoCDay11(unittest.TestCase):
    def test_get_seat(self):
        waiting_area = aoc.WaitingArea([["L", "#"]])
        self.assertEqual("L", waiting_area.get_seat((0, 0)))
        self.assertEqual("#", waiting_area.get_seat((1, 0)))

    def test_count_occupied_neighbors(self):
        waiting_area = aoc.WaitingArea([["L", "#"]])
        self.assertEqual(1, waiting_area.count_occupied_neighbors(0, 0))
        self.assertEqual(0, waiting_area.count_occupied_neighbors(1, 0))

    def test_count_occupied_neighbors_square(self):
        waiting_area = aoc.WaitingArea([["L", "#"], ["#", "#"]])
        self.assertEqual(3, waiting_area.count_occupied_neighbors(0, 0))
        self.assertEqual(2, waiting_area.count_occupied_neighbors(0, 1))

    def test_count_occupied_neighbors_medium(self):
        waiting_area = aoc.WaitingArea(
            [["#", ".", "#"], ["#", "#", "#"], ["#", ".", "#"]]
        )
        self.assertEqual(3, waiting_area.count_occupied_neighbors(0, 1))
        self.assertEqual(6, waiting_area.count_occupied_neighbors(1, 1))

    def test_waiting_area_step(self):
        waiting_area = aoc.WaitingArea(
            [["#", "#", "#"], ["#", "#", "#"], ["#", ".", "#"]]
        )
        self.assertEqual("#", waiting_area.get_seat((1, 0)))
        waiting_area.step()
        self.assertEqual("L", waiting_area.get_seat((1, 0)))

    def test_count_occupied_neighbors_visible_sightline(self):
        waiting_area = aoc.WaitingArea(
            [["#", ".", "#"], ["#", "#", "#"], ["#", ".", "#"]], max_depth=None
        )
        self.assertEqual(3, waiting_area.count_occupied_neighbors(0, 0))

    def test_count_occupied_neighbors_empty_seat_in_sightline(self):
        waiting_area = aoc.WaitingArea([[".", "L", ".", "L", "#"]], max_depth=None)
        self.assertEqual(0, waiting_area.count_occupied_neighbors(0, 0))


class TestAocDay12(unittest.TestCase):
    def test_ferry(self):
        navigation_instructions = ["F10", "N3", "F7", "R90", "F11"]
        ferry = aoc.Ferry()
        ferry.move("F10")
        self.assertEqual((10, 0), ferry.get_coordinates())
        ferry.move("N3")
        self.assertEqual((10, 3), ferry.get_coordinates())
        ferry.move("F7")
        self.assertEqual((17, 3), ferry.get_coordinates())
        ferry.move("R90")
        self.assertEqual(90, ferry.get_bearing())
        self.assertEqual((17, 3), ferry.get_coordinates())
        ferry.move("F11")
        self.assertEqual((17, -8), ferry.get_coordinates())
        coordinates = ferry.get_coordinates()
        manhattan_distance = abs(coordinates[0]) + abs(coordinates[1])
        self.assertEqual(25, manhattan_distance)

    def test_waypoint_ferry(self):
        navigation_instructions = ["F10", "N3", "F7", "R90", "F11"]
        ferry = aoc.WaypointFerry()
        self.assertEqual((10, 1), ferry.get_waypoint())
        ferry.move("F10")
        self.assertEqual((100, 10), ferry.get_coordinates())
        ferry.move("N3")
        self.assertEqual((10, 4), ferry.get_waypoint())
        self.assertEqual((100, 10), ferry.get_coordinates())
        ferry.move("F7")
        self.assertEqual((170, 38), ferry.get_coordinates())
        ferry.move("R90")
        self.assertEqual((4, -10), ferry.get_waypoint())
        self.assertEqual((170, 38), ferry.get_coordinates())
        ferry.move("F11")
        self.assertEqual((214, -72), ferry.get_coordinates())
        coordinates = ferry.get_coordinates()
        manhattan_distance = abs(coordinates[0]) + abs(coordinates[1])
        self.assertEqual(286, manhattan_distance)


if __name__ == "__main__":
    unittest.main()
