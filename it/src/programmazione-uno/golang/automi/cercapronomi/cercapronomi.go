/**
 * Copyright (C) 2023 Giulio Carlo
 *
 * This file is part of programmazione-uno.
 *
 * programmazione-uno is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * programmazione-uno is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with programmazione-uno.  If not, see <http://www.gnu.org/licenses/>.
 */

// E10.2.7, pag. 536
// Da "Foundations of Computer Science: C Edition", Aho, Ullmann
//
// Convertire l'automa dell'esercizio 10.2.6 in una funzione e usarla in un
// programma per trovare tutti i posti dove i pronomi di terza persona singolare
// appaiono come sottostringhe.
package main

import (
	"fmt"
	"io"
	"os"
)

// readFileToBuffer legge il contenuto di un file dato il suo nome e
// restituisce il contenuto come un buffer di byte.
func readFileToBuffer(f string) ([]byte, error) {
	// Apre il file in lettura
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Legge il contenuto del file in un buffer
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// stringsFromBuffer legge i caratteri da un buffer di byte e chiama la funzione
// searchPronoms per determinare la posizione del pronome inglese singolare
// all'interno del file.
func stringsFromBuffer(buffer []byte) {

	// Converte il buffer di byte in una stringa
	txt := string(buffer)

	pos := searchPronoms(txt)

	for _, v := range pos {
		fmt.Println(v)
	}

}

func searchPronoms(t string) (gotcha []int) {

	var (
		q rune
		g int
	)

	for p, r := range t {

		switch q {
		case 0:
			q = gotoQ0(r)
		case 1:
			q = gotoQ1(r)
		case 2:
			q = gotoQ2(r)
		case 3:
			q = gotoQ3(r)
		case 4:
			q = gotoQ4(r)
		case 5:
			q = gotoQ5(r)
		case 6:
			q = gotoQ6(r)
		case 7:
			q = gotoQ7(r)
		case 8:
			q = gotoQ8(r)

		}

		if q == 8 {
			gotcha = append(gotcha, p-g)
			g = 0
		}

		g++

	}

	return
}

func gotoQ0(r rune) rune {
	if r == 'h' {
		// Passa allo stato q1
		return 1
	}

	if r == 's' {
		// passa allo stato q6
		return 6
	}

	// Per ogni altro carattere, compreso lo spazio bianco, resta in q0
	return 0
}

func gotoQ1(r rune) rune {

	if r == 'e' {
		// Passa allo stato q2
		return 2
	}

	if r == 'i' {
		// Passa allo stato q5
	}

	// Per ogni altro carattere, compreso lo spazio bianco, torna in q0
	return 0
}

func gotoQ2(r rune) rune {

	if r == 'r' {
		// Passa allo stato q3
		return 3
	}

	if r == ' ' {
		// Passa allo stato di accettazione q8
		return 8
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func gotoQ3(r rune) rune {

	if r == 's' {
		// Passa allo stato q4
		return 4
	}

	if r == ' ' {
		// Passa allo stato di accettazione q8
		return 8
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func gotoQ4(r rune) rune {

	if r == ' ' {
		// Passa allo stato di accettazione q8
		return 8
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func gotoQ5(r rune) rune {

	if r == 's' {
		// Passa allo stato q4
		return 4
	}

	if r == 'm' {
		// Passa allo stato q7
		return 7
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func gotoQ6(r rune) rune {

	if r == 'h' {
		// Passa allo stato q1
		return 1
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func gotoQ7(r rune) rune {

	if r == ' ' {
		// Passa allo stato di accettazione q8
		return 8
	}

	// Per ogni altro carattere torna in q0
	return 0

}

func gotoQ8(r rune) rune {

	if r == ' ' {
		// Resta nello stato q8
		return 8
	}

	// Per ogni altro carattere torna in q0
	return 0
}

func main() {

	// Verifica se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: cercapronomi <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando
	filename := os.Args[1]

	// Legge il contenuto del file in buffer di byte
	buffer, err := readFileToBuffer(filename)
	if err != nil {
		wrapperErr := fmt.Errorf("error reading file: %w", err)
		fmt.Println(wrapperErr)
	}

	// Legge le parole dal buffer e avvia l'automa per determinare la posizione
	// dei pronomi personali inglesi
	stringsFromBuffer(buffer)
}
