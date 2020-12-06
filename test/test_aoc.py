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


if __name__ == "__main__":
    unittest.main()
