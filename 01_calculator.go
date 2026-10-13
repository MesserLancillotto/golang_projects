package main

import (
	"errors"
	"fmt"
)

type operation func(a, b float64) (float64, error)

var ops = map[string]operation{
	"+": func(a, b float64) (float64, error) { return a + b, nil },
	"-": func(a, b float64) (float64, error) { return a - b, nil },
	"*": func(a, b float64) (float64, error) { return a * b, nil },
	"/": func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	},
}

func main() {
	var opCmd string
	var a, b float64

	for {
		fmt.Print("Operazione (+, -, *, /) o {/quit, /exit} per terminare: ")
		if _, err := fmt.Scan(&opCmd); err != nil {
			fmt.Println("Errore durante la lettura dell'operatore.")
			return
		}

		if opCmd == "/quit" || opCmd == "/exit" {
			break
		}

		fn, exists := ops[opCmd]
		if !exists {
			fmt.Println("Operatore non valido.")
			continue
		}

		fmt.Print("a: ")
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Println("Input non valido per 'a'. Inserisci un valore numerico.")
			clearInputBuffer()
			continue
		}

		fmt.Print("b: ")
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Println("Input non valido per 'b'. Inserisci un valore numerico.")
			clearInputBuffer()
			continue
		}

		res, err := fn(a, b)
		if err != nil {
			fmt.Printf("Errore: %v\n", err)
			continue
		}

		fmt.Printf("-> %g\n", res)
	}
}

func clearInputBuffer() {
	var discard string
	fmt.Scanln(&discard)
}