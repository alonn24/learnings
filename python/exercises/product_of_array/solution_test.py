from product_of_array import Solution
solution = Solution()


def assertCompare(actual, expected):
    assert len(actual) == len(expected)
    assert all([a == b for a, b in zip(actual, expected)])


def testSolution():
    actual = solution.productExceptSelf([1, 2, 3, 4])
    assertCompare(actual, [24, 12, 8, 6])
