def getDigit(s):
    value = ord(s) - ord('0')
    return value if value >= 0 and value <= 9 else None


class Solution:
    def myAtoi(self, s: str) -> int:
        factor = 1
        state = 0
        result = 0
        for c in s:
            digit = getDigit(c)
            if state == 0 and c == ' ':
                continue
            elif state == 0 and c in '-+':
                state = 1
                factor = -1 if c == '-' else 1
            elif state in [0, 1] and digit is not None:
                state = 1
                result = result*10+digit
            else:
                break
        result = result*factor
        return max(-2**31, min(result, 2**31-1))
