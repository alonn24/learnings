from recursive import Solution
solution = Solution()


def test_recursion():
    assert solution.maxCoins([]) == 0
    assert solution.maxCoins([1]) == 1
    assert solution.maxCoins([1, 5]) == 10
    assert solution.maxCoins([5]) == 5
    assert solution.maxCoins([3, 5, 2]) == 39
    assert solution.maxCoins([3, 1, 5, 8]) == 167


def test_timeLimmit():
    assert solution.maxCoins(
        [2, 4, 8, 4, 0, 7, 8, 9, 1, 2, 4, 7, 1, 7, 3, 6]) == 3304
    # assert solution.maxCoins([x for x in range(1000)]) == 10
