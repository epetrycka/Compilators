# Opis tokenów

* Język implementacji: GoLang

## Spis tokenów

| Token          | Reguła                        | Opis                                      |
|----------------|-------------------------------|-------------------------------------------|
| **Liczba_ZP**  | `'cyfra' {'cyfra'}+`           | Liczba całkowita bez znaku                |
| **Identyfikator** | `'Litera' {'litera' \| 'cyfra'}+` | Identyfikator zaczyna się literą      |
| **Dodawanie**  | `'+'`                        | Symbol operacji dodawania algebraicznego  |
| **Odejmowanie** | `'-'`                        | Symbol operacji odejmowania algebraicznego |
| **Mnożenie**   | `'*'`                        | Symbol operacji mnożenia algebraicznego  |
| **Dzielenie**  | `'/'`                        | Symbol operacji dzielenia algebraicznego |
| **Nawias_L**   | `'('`                        | Symbol rozpoczęcia nawiasu               |
| **Nawias_P**   | `')'`                        | Symbol zakończenia nawiasu               |

Komenda na odpalenie skanera:

`python skanerv2.py -file=input.txt`

Program:

- pobiera sekwencyjnie z pliku wyrażenia
- po napotkaniu spacji zamyka aktualne wyrażenie i przekazuje do funkcji _scanner_ wraz z pozycją _position=0_
- funkcja _scanner_ zwraca pierwszy napotkany token zaczynając od _position_ (wypisuje w konsoli parę (klucz, wartość) )
- funkcja _main_ wywołuje _scanner_ z zaktualnioną wartością _position_ do znalezienia kolejnego tokenu aż do końca wyrażenia
- po przeanalizowanym wyrażeniu program pobiera kolejne wyrażenie do następnej napotkanej spacji i wykonuje to samo
- w przypadku znalezienia nieoczekiwanego znaku zwraca informację w konsoli i wypisuje miejsce wystąpienia znaku (row, column)