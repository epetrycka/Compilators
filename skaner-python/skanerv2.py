import argparse
from tokens import whiteSpaces, operators, numbers, letters

def scanner(expression: str, position: int) -> tuple[str, str, int]:
    token = ''
    while (position < len(expression)):
        char = expression[position]
        position += 1

        if char in operators:
            return char, operators[char], position
        
        if char in numbers:
            token += char
            while expression[position] in numbers:
                token += expression[position]
                position += 1
            return token, "Liczba_CBZ", position
        
        if char in letters:
            token += char
            while position < len(expression) and (expression[position] in numbers or expression[position] in letters):
                token += expression[position]
                position += 1
            return token, "Identifier", position

    return token, 0, position

def process_expression(expression: str, position: int, row: int, columns: int) -> int:
    while position < len(expression):
        token, value, position = scanner(expression, position)
        print(f'Token: {token}, value: {value}, row: {row}, column: {position + columns}')
    return position

def main():
    parser = argparse.ArgumentParser(description="Podaj plik wejściowy do skanera")
    parser.add_argument("-file", required=True ,help="Path to file with input data")
    args = parser.parse_args()
    file_name = args.file

    try:
        with open(file_name, "r") as file:
            position = 0
            row = 1
            columns = 0
            expression = ""

            while (char := file.read(1)):
                if char in whiteSpaces:
                    position = process_expression(expression, position, row, columns)
                    columns += len(expression) + 1
                    expression = ""
                    position = 0
                    if char == '\n': 
                        row += 1
                        columns = 0
                else:
                    expression = expression + char

            if char == "" and expression != "":
                _ = process_expression(expression, position, row, columns)

    except Exception as e:
        raise ValueError(e)

if __name__ == "__main__":
    main()