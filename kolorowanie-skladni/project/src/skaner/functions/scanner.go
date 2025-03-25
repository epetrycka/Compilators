package functions

import (
	"fmt"
	"bufio"
	"skaner/tokeny"
)

func Scanner(expression []rune, position int) (string, string, int, bool) {
	//Pobranie znaku z danej pozycji z wyrażenia
	char := expression[position]

	//Jeśli znak jest jednym z operatorów
	if value, exists := tokeny.Operators[char]; exists {
		position += 1
		//Sprawdzenie czy wyrażenie nie jest znakiem początku lub końca komentarza
		if position < len(expression) {
			var token []rune
			token = append(token, char)
			token = append(token, expression[position])
			position += 1
			if token_id, exist := tokeny.LogicExp[string(token)]; exist {
				return string(token), token_id, position, true
			} else {
				position -= 1
				return string(char), value, position, true
			}
		}
		return string(char), value, position, true
	}

	//Sprawdzenie czy twyrażenie nie jest nawiasem
	if value, exists := tokeny.Brackets[char]; exists {
		position += 1
		return string(char), value, position, true
	}

	//Sprawdzenie czy wyrażenie nie jest wartością liczbową
	if tokeny.Numbers[char] {
		var token []rune
		token = append(token, char)
		position += 1
		//Pobranie całej liczby z wyrażenia i zwrócenie jej 
		for position < len(expression) && tokeny.Numbers[expression[position]]{
			token = append(token, expression[position])
			position += 1
		}
		return string(token), "number", position, true 
	}

	//Sprawdzenie czy wyrażenie nie jest identyfikatorem zaczynającym się literą
	if tokeny.Letters[char] {
		var token []rune
		token = append(token, char)
		position += 1
		//Pobranie całego identyfikatora literowo-liczbowego i zwrócenie go
		for position < len(expression) && (tokeny.Numbers[expression[position]] || tokeny.Letters[expression[position]]){
			token = append(token, expression[position])
			position += 1
		}
		return string(token), string(token), position, true 
	}

	return string(char), "invalid", position, false
}

func Process_expression(expression []rune, position int, row int, column int, writer *bufio.Writer, comment_line bool, comment_block bool, text bool) (int, bool, bool, bool) {
	//Wywoływanie funkcji skaner aż do końca wyrażenia
	for position < len(expression){
		//Wywołanie skanera
		token, id_tokenu, newPosition, err :=  Scanner(expression, position)
		//Pobranie nowej pozycji w wyrażeniu
		position = newPosition
		//Badanie wystąpienia nieoczekiwanego znaku w wyrażeniu
		if !err {
			fmt.Printf("Error: Unexpected token: %s, in row: %d, column: %d \n", token, row, (position + column + 1 - len(token)))
			return position, comment_line, comment_block, text
		}
		//Wyświetlenie w konsoli
		fmt.Printf("Token: %s, id_tokenu: %s, row: %d, column: %d \n", token, id_tokenu, row, (position + column - len(token)))
		
		//Zapis do pliku elementów html
		//Zaktualizowanie zmiennej występowania komentarza wieloliniowego
		if id_tokenu == "comment_block_start"{
			comment_block = true
		}
		if id_tokenu == "comment_line" {
			comment_line = true
		}
		if id_tokenu == "comment_block_end" {
			comment_block = false
		}
		//Aktualizacja wartości tokenów, aby przypisać atrybut komentarza dla komentarzy
		if comment_line || comment_block {
			id_tokenu = "comment_line"
		}
		if id_tokenu == "double_quotation" && !text{
			text = true
		} else if id_tokenu == "double_quotation" && text {
			text = false
		}
		if text {
			id_tokenu = "text"
		}

		//Sprawdzenie nazwy klasy w słowniku { id_tokenu : class_name }
		if class, exist := tokeny.KeyWords[id_tokenu]; exist{
			writer.WriteString("<span class=\"" + class + "\">" + token + "</span>")
			writer.Flush()
		} else {
			//Jeśli id_tokenu nie występuje w pliku zapisujemy bez klasy
			writer.WriteString(token)
			writer.Flush()
		}
	}
	return position, comment_line, comment_block, text
}