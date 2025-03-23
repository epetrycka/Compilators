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
        }
        .if { color: #ff79c6; }   /* Różowy */
        .else { color: #00a2ff; }   /* Niebieski */
        .for { color: #0008fb; }   /* Granatowy */
        .return { color: #ff0000; }   /* Czerwony */
        .func { color: #fbff00; }   /* Żółty */
        .var { color: #ff7700; }   /* Pomarańczowy */
        .type { color: #a870f5; }      /* Fioletowy */
        .number { color: #bfbfbf; }    /* Szary */
        .operator { color: #615243; }  /* Brązowy */
        .comment { color: #50fa7b; }   /* Zielony */
        .brackets { color: #ffffff; }  /* Biały */
        .seperators { color: #faffe8; }  /* Biały */
        .logic { color: #00fffb; }  /* Turkusowy */
        .compare { color: #ff0090; } /* Malinowy */
        .print { color: #b3ff00; } /* Seledynowy */
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

var Operators = map[rune]string{
	'+' : "Dodawanie",
	'-': "Odejmowanie",
	'*': "Mnozenie",
	'/': "Dzielenie",
	'=': "Przypisanie",
	'%': "Modulo",
	'<': "Mniejsze",
	'>': "Wieksze",
	',': "Przecinek",
	';': "Średnik",
	':': "Dwukropek",
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
	"+" : "operator",
	"-" : "operator",
	"*" : "operator",
	"/" : "operator",
	"%" : "operator",
	"=" : "operator",
	"&&" : "logic",
	"||" : "logic",
	"!" : "logic",
	"==" : "compare",
	"!=" : "compare",
	">" : "compare",
	"<" : "compare",
	">=" : "compare",
	"<=" : "compare",
	"(" : "brackets",
	")" : "brackets",
	"{" : "brackets",
	"}" : "brackets",
	"," : "seperators",
	";" : "seperators",
	":" : "seperators",
}