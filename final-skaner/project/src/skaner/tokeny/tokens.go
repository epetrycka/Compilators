package tokeny

var WhiteSpaces = map[rune]bool{
	' ':  true,
	'\n': true,
	'\r': true,
	'\t': true,
	'\v': true,
	'\f': true,
}

var Operators = map[rune]string{
	'+' : "Dodawanie",
	'-': "Odejmowanie",
	'*': "Mnozenie",
	'/': "Dzielenie",
	'(': "Nawias_L",
	')': "Nawias_P",
}

var Numbers = map[rune]bool{
	'0' : true,
	'1' : true,
	'2' : true,
	'3' : true,
	'4' : true,
	'5' : true,
	'6' : true,
	'7' : true,
	'8' : true,
	'9' : true,
}

var Letters = map[rune]bool{
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true,
	'g': true, 'h': true, 'i': true, 'j': true, 'k': true, 'l': true,
	'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true,
	's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true,
	'y': true, 'z': true,

	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true,
	'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true,
	'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true,
	'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true,
	'Y': true, 'Z': true,
}