package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/tokeny"
)

func scanner(expression []rune, position int) (string, string, int, bool) {
	char := expression[position]

	if value, exists := tokeny.Operators[char]; exists {
		position += 1
		return string(char), value, position, true
	}

	if tokeny.Numbers[char] {
		var token []rune
		token = append(token, char)
		position += 1
		for position < len(expression) && tokeny.Numbers[expression[position]]{
			token = append(token, expression[position])
			position += 1
		}
		return string(token), "Liczba", position, true 
	}

	if tokeny.Letters[char] {
		var token []rune
		token = append(token, char)
		position += 1
		for position < len(expression) && (tokeny.Numbers[expression[position]] || tokeny.Letters[expression[position]]){
			token = append(token, expression[position])
			position += 1
		}
		return string(token), "Identyfikator", position, true 
	}

	return string(char), "invalid", position, false
}

func process_expression(expression []rune, position int, row int, column int) int {
	for position < len(expression){
		token, value, newPosition, err := scanner(expression, position)
		position = newPosition
		if !err {
			fmt.Printf("Error: Unexpected token: %s, in row: %d, column: %d \n", token, row, (position + column + 1 - len(token)))
			return position
		}
		fmt.Printf("Token: %s, value: %s, row: %d, column: %d \n", token, value, row, (position + column - len(token)))
	}
	return position
}

func main(){
	file_name := flag.String("file", "", "File name with input data")
	flag.Parse()

	if *file_name == "" {
		fmt.Println(`Error: Flag -file= is required.
					Template: go run main.go -file=input.txt`)
		os.Exit(1)
	}

	file, err := os.Open(*file_name)
	if err != nil{
		fmt.Printf("Error: Unable to open file: %s", *file_name)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Printf("Scanner started in file: %s \n", *file_name)
	var expression []rune
	position := 0
	row := 1
	columns := 1

	for {
		char, _, err := reader.ReadRune()
		
		if err != nil{
			break
		}

		// fmt.Printf("nowy znak %q \n", char)
		if tokeny.WhiteSpaces[char] {
			if char == '\r' { continue }
			newPosition := process_expression(expression, position, row, columns)
			
			if newPosition != len(expression){
				os.Exit(1)
			}
			
			columns += len(expression) + 1
			expression = expression[:0]
			position = 0

			if char == '\n' {
				row += 1
				columns = 1
			}
		} else {
			expression = append(expression, char)
		}
	}

	newPosition := process_expression(expression, position, row, columns)

	if newPosition != len(expression){
		os.Exit(1)
	}
}