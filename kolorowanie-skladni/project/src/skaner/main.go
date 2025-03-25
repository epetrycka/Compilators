package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/tokeny"
)

func scanner(expression []rune, position int) (string, string, int, bool) {
	//Pobranie znaku z danej pozycji z wyrażenia
	char := expression[position]

	//Jeśli znak jest jednym z operatorów
	if value, exists := tokeny.Operators[char]; exists {
		position += 1
		//Sprawdzenie czy wyrażenie nie jest znakiem początku lub końca komentarza
		if position < len(expression) && expression[position] == '/' && string(char) == "/"{
			var token []rune
			token = append(token, char)
			token = append(token, expression[position])
			position += 1
			return string(token), "comment", position, true
		}
		if position < len(expression) && expression[position] == '*' && string(char) == "/"{
			var token []rune
			token = append(token, char)
			token = append(token, expression[position])
			position += 1
			return string(token), "comment", position, true
		}
		if position < len(expression) && expression[position] == '/' && string(char) == "*"{
			var token []rune
			token = append(token, char)
			token = append(token, expression[position])
			position += 1
			return string(token), "comment", position, true
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

func process_expression(expression []rune, position int, row int, column int, writer *bufio.Writer, comment_line bool, comment_block bool) (int, bool, bool) {
	//Wywoływanie funkcji skaner aż do końca wyrażenia
	for position < len(expression){
		//Wywołanie skanera
		token, id_tokenu, newPosition, err := scanner(expression, position)
		//Pobranie nowej pozycji w wyrażeniu
		position = newPosition
		//Badanie wystąpienia nieoczekiwanego znaku w wyrażeniu
		if !err {
			fmt.Printf("Error: Unexpected token: %s, in row: %d, column: %d \n", token, row, (position + column + 1 - len(token)))
			return position, comment_line, comment_block
		}
		//Wyświetlenie w konsoli
		fmt.Printf("Token: %s, id_tokenu: %s, row: %d, column: %d \n", token, id_tokenu, row, (position + column - len(token)))
		
		//Zapis do pliku elementów html
		//Zaktualizowanie zmiennej występowania komentarza wieloliniowego
		if id_tokenu == "comment"{
			if token == "/*" {
				comment_block = true
			}
			if token == "//" {
				comment_line = true
			}
			if token == "*/" {
				comment_block = false
			}
		}
		//Aktualizacja wartości tokenów, aby przypisać atrybut komentarza dla komentarzy
		if comment_line || comment_block {
			id_tokenu = "comment"
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
	return position, comment_line, comment_block
}

func main(){
	//Identyfikacja flagi
	file_name := flag.String("file", "", "File name with input data")
	flag.Parse()
	if *file_name == "" {
		fmt.Println(`Error: Flag -file= is required.
					Template: go run main.go -file=input.txt`)
		os.Exit(1)
	}

	//Otwieranie pliku wejściowego
	file, err := os.Open(*file_name)
	if err != nil{
		fmt.Printf("Error: Unable to open file: %s", *file_name)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	//Otwieranie pliku wyjściowego
	output_file := "output.html"
	output, err := os.Create(output_file)
	if err != nil {
		fmt.Println("Błąd przy tworzeniu pliku:", err)
		return
	}
	defer output.Close()
	writer := bufio.NewWriter(output)

	//Zapis pierwszych linii html
	writer.WriteString(tokeny.Inital)
	writer.Flush()

	fmt.Printf("Scanner started in file: %s \n", *file_name)
	//Deklaracja niezbędnych zmiennych
	var expression []rune
	position := 0
	row := 1
	columns := 1
	comment_line := false
	comment_block := false
	
	//Wywoływanie skanera w pętli do ostatniego wyrażenia
	for {
		//Wczytujemy jeden znak z pliku
		char, _, err := reader.ReadRune()
		
		//Sprawdzamy czy nie wystąpił błąd podczas czytania pliku
		if err != nil{
			break
		}

		// fmt.Printf("nowy znak %q \n", char)
		//Operacje wykonywane po zakończeniu wczytywania wyrażenia do pierwszej spacji
		if tokeny.WhiteSpaces[char] {
			if char == '\r' { continue } //Znak końca linii - pomijamy
			//Wywołanie funkcji procesowania i zapis zmiennych wartości o wystąpieniu znaku komentarza
			newPosition, newComment, newCommentBlock := process_expression(expression, position, row, columns, writer, comment_line, comment_block)
			comment_line = newComment
			comment_block = newCommentBlock

			//Zapisujemy białe znaki do pliku
			writer.WriteString(string(char))
			writer.Flush()

			//Sprawdzenie wystąpienia błędu podczas wywoływania skanera
			if newPosition != len(expression){
				os.Exit(1)
			}
			
			//Czyszczenie pamięci na nowe wyrażenie, ustawienie pointera na 0, liczenie kolumn w danym wierszu
			columns += len(expression) + 1
			expression = expression[:0]
			position = 0

			//Dla nowej linii obliczamy wiersz, zerujemy kolumnę i kończymy komentarz liniowy
			if char == '\n' {
				row += 1
				columns = 1
				comment_line = false
			}
		//Jeśli nie napotkamy spacji tworzymy wyrażenie
		} else {
			expression = append(expression, char)
		}
	}

	//Wywołanie skanera jeszcze raz na końcu
	newPosition, _, _ := process_expression(expression, position, row, columns, writer, comment_line, comment_block)
	//Badanie błędów skanera i procesu skanowania
	if newPosition != len(expression){
		os.Exit(1)
	}

	//Zapis końcowych linii html
	writer.WriteString(tokeny.End)
	writer.Flush()
}