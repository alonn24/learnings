from spiral_matrix import Solution
solution = Solution()


def assertCompare(actual, expected):
    assert len(actual) == len(expected)
    assert all([a == b for a, b in zip(actual, expected)])


def testSpiralMatrix():
    assertCompare(solution.spiralOrder([[]]), [])
    assertCompare(solution.spiralOrder([[1]]), [1])
    assertCompare(solution.spiralOrder([[1, 2]]), [1, 2])
    assertCompare(solution.spiralOrder([[1], [2]]), [1, 2])
    actual = solution.spiralOrder(
        [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]])
    assertCompare(actual, [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7])


def testSubmit():
    assertCompare(solution.spiralOrder([[1], [2], [3], [4], [5], [6], [
                  7], [8], [9], [10]]), [x for x in range(1, 11)])
    assertCompare(solution.spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [
                  13, 14, 15, 16]]), [1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10])
