package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/tokeny"
)

func Scanner(file_name string){
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error: Unable to open file:", file_name)
		os.Exit(1)
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
			token := scanNumber(reader, ch)
			fmt.Println("Liczba : ", token)
			continue
		}

		if tokeny.Letters[ch]{
			token := scanIdentifier(reader, ch)
			fmt.Println("Identyfikator : ", token)
			continue
		}

		fmt.Printf("Unexpected token in line %d, column %d: %s\n", row, column, string(ch))
	}
}

func scanNumber(reader *bufio.Reader, first rune) string {
	number := string(first)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if tokeny.Numbers[ch] {
			number += string(ch)
		} else {
			_ = reader.UnreadRune()
			break
		}
	}
	return number
}

func scanIdentifier(reader *bufio.Reader, first rune) string{
	id := string(first)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if tokeny.Letters[ch] {
			id += string(ch)
		} else if tokeny.Numbers[ch] {
			id += string(ch)
		} else {
			_ = reader.UnreadRune()
			break
		}
	}
	return id
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

	Scanner(*file_name)

	fmt.Println("\nScanner finished its job.")
	os.Exit(0)
}