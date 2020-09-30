package zigzag

import (
	strs "strings"
)

func inc(n int) int {
	return n + 1
}

func dec(n int) int {
	return n - 1
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([]strs.Builder, numRows)
	chars := []rune(s)
	currentRow := 0
	incdec := inc
	for i := 0; i < len(chars); i++ {
		sb := &rows[currentRow]
		if currentRow == 0 {
			incdec = inc
		} else if currentRow == numRows-1 {
			incdec = dec
		}
		currentRow = incdec(currentRow)
		sb.WriteRune(chars[i])
	}
	mainSb := strs.Builder{}
	for sb := range rows {
		mainSb.WriteString(rows[sb].String())
	}
	return mainSb.String()
}
