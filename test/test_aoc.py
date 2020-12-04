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


if __name__ == "__main__":
    unittest.main()
