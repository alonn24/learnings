package validnumber

func isIn(c rune, arr []rune) bool {
	for _, r := range arr {
		if r == c {
			return true
		}
	}
	return false
}

func takeWhile(s string, c []rune) (string, string) {
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		if !isIn(chars[i], c) {
			return string(chars[0:i]), string(chars[i:len(chars)])
		}
	}
	return s, ""
}

func mustHave(s1 string, s2 string) bool {
	if len(s1) > 0 {
		return len(s2) > 0
	}
	return true
}

func isNumber(s string) bool {
	_, rest := takeWhile(s, []rune{' '})
	plusMinus, rest := takeWhile(rest, []rune{'-', '+'})
	number, rest := takeWhile(rest, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	dot, rest := takeWhile(rest, []rune{'.'})
	friction, rest := takeWhile(rest, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	exponent, rest := takeWhile(rest, []rune{'e'})
	exponentPlusMinus, rest := takeWhile(rest, []rune{'-', '+'})
	exponentNumber, rest := takeWhile(rest, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	_, rest = takeWhile(rest, []rune{' '})

	res := (len(plusMinus) == 0 || plusMinus == "+" || plusMinus == "-") &&
		(len(number) > 0 || len(friction) > 0) &&
		(len(dot) == 0 || dot == ".") &&
		(len(exponent) == 0 || exponent == "e") &&
		len(rest) == 0 &&
		mustHave(exponent, exponentNumber) &&
		mustHave(exponentPlusMinus, exponent)
	return res
}
