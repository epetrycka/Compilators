import argparse
from tokens import whiteSpaces, operators, numbers, letters

def scanner(expression: str, position: int, row: int, column: int) -> tuple[str, str, int]:
    token = ''
    char = expression[position]
    position += 1

    if char in operators:
        return char, operators[char], position

    if char in numbers:
        token += char
        while position < len(expression) and expression[position] in numbers:
            token += expression[position]
            position += 1
        return token, "Liczba_CBZ", position

    if char in letters:
        token += char
        while position < len(expression) and (expression[position] in numbers or expression[position] in letters):
            token += expression[position]
            position += 1
        return token, "Identifier", position

    raise ValueError(f"Nieoczekiwany znak '{char}' w wierszu {row}, kolumna {column + 1}")

def process_expression(expression: str, position: int, row: int, columns: int):
    while position < len(expression):
        try:
            token, value, position = scanner(expression, position, row, columns + position)
            print(f'Token: {token}, value: {value}, row: {row}, column: {columns + position - len(token)}')
        except ValueError as e:
            print(f"Błąd: {e}")
            return

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
                    process_expression(expression, position, row, columns)
                    columns += len(expression) + 1
                    expression = ""
                    position = 0
                    if char == '\n': 
                        row += 1
                        columns = 0
                else:
                    expression += char

            if expression:
                process_expression(expression, position, row, columns)

    except Exception as e:
        print(f"Wystąpił nieoczekiwany błąd: {e}")

if __name__ == "__main__":
    main()
