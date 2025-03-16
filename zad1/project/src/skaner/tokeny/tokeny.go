package tokeny

import "regexp"

type Token struct{
	Regex *regexp.Regexp
	Nazwa string
}

var Tokeny = []Token{
	{regexp.MustCompile(`^\+`), "Dodawanie"},
	{regexp.MustCompile(`^-`), "Odejmowanie"},
	{regexp.MustCompile(`^\*`), "Mnozenie"},
	{regexp.MustCompile(`^/`), "Dzielenie"},
	{regexp.MustCompile(`^\(`), "Nawias_L"},
	{regexp.MustCompile(`^\)`), "Nawias_P"},
	{regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*`), "Identyfikator"},
	{regexp.MustCompile(`^[0-9]+`), "Liczba calkowita"},
}