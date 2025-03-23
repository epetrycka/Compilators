# Kolorowanie składni języka
## Zakres
* przygotowanie własnego spisu tokenów do kolorowania
* opracowanie diagramu przejść do rozpoznawania tokenów
* implementacja skanera z kolorowanem składni wejściowego formatu
* informacje wstępne o projektach
## Zadania
* Przygotuj tabelę tokenów własnego formatu.
* Bazując na przygotowanym własnym formacie np. na uproszczonym języku programowania, napisz program kolorujący składnię.
* Zaimplementuj i wykorzystaj skaner rozpoznający tokeny wg przygotowanej tabeli tokenów. 
* Program powinien wczytywać z pliku kod w wybranym formacie i zwrócić do innego pliku pokolorowany kod. 
* Jako formatu wyjściowego można wykorzystać np. HTML. W pliku wyjściowym powinien być zachowany układ tekstu z pliku wejściowego.

### Spis tokenów i kolorów

| Token       | Znaczenie             | Kolor        |
|-------------|-----------------------|--------------|
| `if`        | Warunek               | Różowy       |
| `else`      | Inaczej               | Niebieski    |
| `for`       | Pętla                 | Granatowy      |
| `return`    | Zwracanie wartości    | Czerwony     |
| `func`      | Deklaracja funkcji    | Żółty       |
| `var`       | Deklaracja zmiennej   | Pomarańczowy |
| `string`    | Typ tekstowy          | Fioletowy    |
| `int`       | Typ całkowity         | Fioletowy    |
| `print`     | Drukowanie wyniku     | Seledynowy    |
| `+ - * / %` | Operatory matematyczne | Brązowy      |
| `&& \|\| !`   | Operatory logiczne    | Turkusowy    |
| `== != < > <= >=` | Operatory porównania | Malinowy  |
| `123, 3.14` | Liczby                | Szary        |
| `//`        | Komentarz jednoliniowy | Zielony      |
| `/* ... */` | Komentarz wieloliniowy | Zielony      |
| `(` `)` `{` `}` `[ ]` | Nawiasy i znaki specjalne | Biały |
| `, ;`       | Separatory            | Biały       |
