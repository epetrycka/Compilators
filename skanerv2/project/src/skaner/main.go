package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/tokeny"
)

type ScannerError struct {
	Row    int
	Column int
	Char   rune
}

func (e ScannerError) Error() string {
	return fmt.Sprintf("Unexpected token in line %d, column %d: %s", e.Row, e.Column, string(e.Char))
}

func Scanner(file_name string) error{
	file, err := os.Open(file_name)
	if err != nil {
		return fmt.Errorf("Error: Unable to open file: %s", file_name)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	row := 1
	column := 0
	for {
		column += 1
		ch, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		if tokeny.WhiteSymbols[ch]{
			if ch == '\n'{
				row += 1
				column = 0
			}
			continue
		}

		if name, found := tokeny.Tokeny[ch];  found{
			fmt.Println(name, " : ", string(ch))
			continue
		}

		if tokeny.Numbers[ch]{
			token, newColumn := scanNumber(reader, ch, column)
			column = newColumn
			fmt.Println("Liczba : ", token)
			continue
		}

		if tokeny.Letters[ch]{
			token, newColumn := scanIdentifier(reader, ch, column)
			column = newColumn
			fmt.Println("Identyfikator : ", token)
			continue
		}

		return ScannerError{Row: row, Column: column, Char: ch}
	}

	return nil
}

func scanNumber(reader *bufio.Reader, first rune, column int) (string, int) {
	number := string(first)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if tokeny.Numbers[ch] {
			number += string(ch)
			column += 1
		} else {
			_ = reader.UnreadRune()
			break
		}
	}
	return number, column
}

func scanIdentifier(reader *bufio.Reader, first rune, column int) (string, int) {
	id := string(first)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if tokeny.Letters[ch] {
			id += string(ch)
			column += 1
		} else if tokeny.Numbers[ch] {
			id += string(ch)
			column += 1
		} else {
			_ = reader.UnreadRune()
			break
		}
	}
	return id, column
}

func main(){
	file_name := flag.String("file", "", "Input file name")
	flag.Parse()

	if *file_name == "" {
		fmt.Println("Error: Flag -file= is required")
		fmt.Println("Template: go run main.go -file=input.txt")
		os.Exit(1)
	}

	fmt.Printf("Scanner started in file: %s \n\n", *file_name)

	err := Scanner(*file_name)
	if err != nil {
		fmt.Println("Scanner Error:", err)
		os.Exit(1)
	}

	fmt.Println("\nScanner finished its job.")
	os.Exit(0)
}