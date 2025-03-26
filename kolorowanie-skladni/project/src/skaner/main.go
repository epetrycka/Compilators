package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/functions"
	"skaner/tokeny"
)

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
	text := false
	
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
			//Wywołanie funkcji procesowania i zapis zmiennych wartości o wystąpieniu znaku komentarza lub tekstu
			newPosition, newComment, newCommentBlock, newText := functions.Process_expression(expression, position, row, columns, writer, comment_line, comment_block, text)
			comment_line = newComment
			comment_block = newCommentBlock
			text = newText

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
	newPosition, _, _, _ := functions.Process_expression(expression, position, row, columns, writer, comment_line, comment_block, text)
	//Badanie błędów skanera i procesu skanowania
	if newPosition != len(expression){
		os.Exit(1)
	}

	//Zapis końcowych linii html
	writer.WriteString(tokeny.End)
	writer.Flush()
}