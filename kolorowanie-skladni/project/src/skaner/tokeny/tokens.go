package tokeny

var Inital = `<!DOCTYPE html>
<html lang="pl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Code</title>
    <style>
        body {
            background-color: #1e1e1e;
            color: white;
            font-family: Consolas, monospace;
            font-size: large;
        }
        .if { color:rgb(255, 0, 144); }   /* Malinowy */
        .else { color: #00a2ff; }   /* Niebieski */
        .for { color:rgb(159, 0, 251); }   /* Fioletowy */
        .return { color:rgb(0, 47, 255); }   /* Granatowy */
        .func { color:rgb(255, 0, 217); }   /* Różowy */
        .var { color: #ff7700; }   /* Pomarańczowy */
        .type { color:rgb(255, 242, 0); }      /* Żółty */
        .number { color:rgb(115, 255, 0); }    /* Seledynowy */
        .operator { color:rgb(145, 108, 71); }  /* Brązowy */
        .comment { color:rgb(0, 188, 25); }   /* Zielony */
        .brackets { color:rgb(255, 0, 0); }  /* Czerowny */
        .seperators { color: #faffe8; }  /* Biały */
        .logic { color:rgb(0, 238, 255); }  /* Turkusowy */
        .compare { color:rgb(255, 225, 0); } /* Seledynowy */
        .print { color:rgb(0, 174, 255); } /* Niebieski */
    </style>
</head>
<body>
<pre>
`
var End = `
</pre>
</body>
</html>
` 

var WhiteSpaces = map[rune]bool{
	' ':  true,
	'\n': true,
	'\r': true,
	'\t': true,
	'\v': true,
	'\f': true,
}

// var Operators = map[rune]string{
// 	'+' : "Dodawanie",
// 	'-': "Odejmowanie",
// 	'*': "Mnozenie",
// 	'/': "Dzielenie",
// 	'=': "Przypisanie",
// 	'%': "Modulo",
// 	'<': "Mniejsze",
// 	'>': "Wieksze",
// 	',': "Przecinek",
// 	';': "Średnik",
// 	':': "Dwukropek",
// 	'!': "Zaprzeczenie",
// 	'&': "Koniunkcja",
// 	'|': "Alternatywa",
// }

var Operators = map[rune]string{
	'+': "plus",
	'-': "minus",
	'*': "multiply",
	'/': "divide",
	'=': "operator",
	'%': "operator",
	'<': "compare",
	'>': "compare",
	',': "operator",
	';': "operator",
	':': "operator",
	'!': "operator",
	'&': "logic",
	'|': "logic",
}

var LogicExp = map[string]string {
	"&&" : "logic",
	"||" : "logic",
	"!" : "logic",
	"==" : "compare",
	"!=" : "compare",
}

var Brackets = map[rune]string{
	'(': "Nawias_L",
	')': "Nawias_P",
	'{': "Klamra_L",
	'}': "Klamra_P",
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

var KeyWords = map[string]string{
	"if" : "if",
	"else" : "else",
	"for" : "for",
	"return" : "return",
	"func" : "func",
	"var" : "var",
	"string" : "type",
	"int" : "type",
	"print" : "print",
	"number" : "number",
	"operator" : "operator",
	"plus" : "operator",
	"minus" : "operator",
	"multiply" : "operator",
	"divide" : "operator",
	"comment" : "comment",
	"compare" : "compare",
	"Nawias_L" : "brackets",
	"Nawias_P" : "brackets",
	"Klamra_L" : "brackets",
	"Klamra_P" : "brackets",
}