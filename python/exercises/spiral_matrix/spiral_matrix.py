import math
from typing import List


def getFrame(matrix, i):
    startRowIndex = i
    endRowIndex = len(matrix) - 1 - i
    startColumnIndex = i
    endColumnIndex = len(matrix[0]) - 1 - i
    hasFrame = startColumnIndex < endColumnIndex and startRowIndex < endRowIndex

    # first row include last item
    firstRow = [matrix[startRowIndex][x]
                for x in range(startColumnIndex, endColumnIndex + 1)]
    # last column exclude first include last
    lastColumn = [matrix[x][endColumnIndex]
                  for x in range(startRowIndex + 1, endRowIndex + 1)]

    if not hasFrame:
        return firstRow + lastColumn
    # last Row exclude first include last
    lastRow = [matrix[endRowIndex][endColumnIndex - x]
               for x in range(1, endColumnIndex - startColumnIndex + 1)]
    # first column exclude first exclude last
    firstColumn = [matrix[endRowIndex-x][startRowIndex]
                   for x in range(1, endRowIndex - startRowIndex)]

    return firstRow+lastColumn+lastRow+firstColumn


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        half = math.ceil(min(len(matrix), len(matrix[0]))/2)
        frames = [getFrame(matrix, i) for i in range(0, half)]
        return [item for sublist in frames for item in sublist]
