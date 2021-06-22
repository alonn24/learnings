from string_to_integer import Solution
solution = Solution()


def test_stringToInteger():
    assert solution.myAtoi("42") == 42
    assert solution.myAtoi("  42") == 42
    assert solution.myAtoi("  -42") == -42
    assert solution.myAtoi("  -    42") == 0
    assert solution.myAtoi("some words 42") == 0
    assert solution.myAtoi("   -42 some words") == -42
    assert solution.myAtoi("-91283472332") == -2147483648
    assert solution.myAtoi("91283472332") == 2147483647


def test_stringToIntegerAdvanced():
    assert solution.myAtoi("00000-42a1234") == 0
