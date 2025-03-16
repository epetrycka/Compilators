package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"skaner/tokeny"
	"strings"
	"unicode"
)

func Scanner(scanner *bufio.Scanner){
	row := 0
	for scanner.Scan(){
		line := scanner.Text()
		row += 1
		
		line = strings.TrimSpace(line)

		if line == ""{
			continue
		}

		inital_size := len(line)

		for len(line) > 0 {
			line = strings.TrimLeftFunc(line, unicode.IsSpace)

			matchfound := false
			for _, token := range tokeny.Tokeny{
				index := token.Regex.FindStringIndex(line)
				if index != nil && index[0] == 0{
					matched := line[:index[1]]
					fmt.Printf("Token: %s, %s\n", token.Nazwa, matched)
					line = line[index[1]:]
					matchfound = true
					break
				}
			}

			if !matchfound {
				fmt.Printf("Error: Token not recognized in row %d, column %d: %s", row, (inital_size - len(line) + 1), string(line[0]))
				line = line[1:]
			}
		}
	}
}

func main(){
	file_name := flag.String("file", "", "Input file name")
	flag.Parse()

	if *file_name == "" {
		fmt.Println("Error: Flag -file= is required")
		fmt.Println("Template: go run main.go -file=input.txt")
		os.Exit(1)
	}

	file, err := os.Open(*file_name)
	if err != nil {
		fmt.Println("Error: Unable to open file: ", *file_name)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("Scanner started in file: %s \n\n", *file_name)

	scanner := bufio.NewScanner(file)

	Scanner(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: Unexpected behaviour during reading file process: ", *file_name)
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nScanner finished its job.")
	os.Exit(0)
}