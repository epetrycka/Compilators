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
	'+': "plus",
	'-': "minus",
	'*': "multiply",
	'/': "divide",
	'=': "assign",
	'%': "modulo",
	'<': "less",
	'>': "more",
	',': "comma",
	';': "semicolon",
	':': "colon",
	'!': "negation",
	'\'' : "single_quotation",
	'"' : "double_quotation",
	'&' : "half_conjuction",
	'|' : "half_alternative",
}

var Brackets = map[rune]string{
	'(': "Param_L",
	')': "Param_P",
	'{': "Curly_L",
	'}': "Curly_P",
}

var LogicExp = map[string]string {
	"&&" : "conjuction",
	"||" : "alternative",
	":=" : "colon_assign",
	"==" : "comparison",
	"!=" : "neg_comparison",
	"<=" : "less_equal",
	">=" : "more_equal",
	"++" : "increase",
	"--" : "decrease",
	"//" : "comment_line",
	"/*" : "comment_block_start",
	"*/" : "comment_block_end",
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
	"plus" : "operator",
	"minus" : "operator",
	"multiply" : "operator",
	"divide" : "operator",
	"assign" : "operator",
	"colon_assign" : "operator",
	"modulo" : "operator",
	"less" : "comparison",
	"more" : "comparison",
	"comma" : "operator",
	"semicolon" : "operator",
    "colon" : "operator",
	"negation" : "operator",
	"increase" : "operator",
	"decrease" : "operator",
	"conjuction" : "logic",
	"alternative" : "logic",
	"comparison" : "comparison",
	"neg_comparison" : "comparison",
	"less_equal" : "comparison",
	"more_equal" : "comparison",
	"Param_L" : "brackets",
	"Param_P" : "brackets",
	"Curly_L" : "brackets",
	"Curly_P" : "brackets",
	"single_quotation" : "text",
	"double_quotation" : "text",
	"text" : "text",
	"comment_line" : "comment",
	"comment_block_start" : "comment",
	"comment_block_end" : "comment",
	"number" : "number",
}

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
        .comparison { color:rgb(255, 225, 0); } /* Seledynowy */
        .print { color:rgb(0, 174, 255); } /* Niebieski */
		.text { color:rgb(6, 194, 34); } /* Zielony */
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