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
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// readFileToBuffer legge il contenuto di un file dato il suo nome
// e restituisce il contenuto come un buffer di byte.
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

// stringsFromBuffer legge le parole da un buffer di byte e chiama la funzione
// searchPattern per determinare se ogni parola è un identificatore legale del
// linguaggio C (lettere seguite da lettere, cifre, oppure underscore) seguito
// da uno spazio bianco.
func stringsFromBuffer(buffer []byte) (val, nval int) {
	// Creazione di un nuovo reader basato sul buffer
	reader := bytes.NewReader(buffer)

	// Creazione di un nuovo scanner che leggerà dal reader
	scanner := bufio.NewScanner(reader)

	// Impostazione dello scanner per scansionare le parole
	scanner.Split(bufio.ScanWords)

	// Ciclo finché ci sono parole da leggere
	for scanner.Scan() {
		// Legge la parola dallo scanner
		word := scanner.Text()

		// Chiama la funzione searchPattern per deteterminare se la parola
		// è un identificatore valido del linguaggio C. In tal caso
		// incrementa il numero di parole valide, altrimenti incrementa
		// il contatore delle parole non valide.
		if searchPattern(word) {
			fmt.Println(word, "OK")
			val++
		} else {
			nval++
			fmt.Println(word, "NO")
		}
	}

	return val, nval
}

func searchPattern(w string) bool {
	// Stato corrente dell'automa
	var q rune

	for _, r := range w {

		switch q {
		// Passa dallo stato q0 allo stato q restituito da gotoQ0
		case 0:
			q = gotoQ0(r)
		// Passa dallo stato q1 allo stato q restituito da gotoQ1
		case 1:
			q = gotoQ1(r)
		// Passa dallo stato q2 allo stato q restituito da gotoQ2
		case 2:
			q = gotoQ2(r)
		// Passa dallo stato q3 allo stato q restituito da gotoQ3
		case 3:
			q = gotoQ3(r)
		}
	}

	return q == 1 || q == 3
}

// qotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
func gotoQ0(r rune) rune {
	// Se la runa in input è una lettera dell'alfabeto vai allo stato di
	// accettazione q1
	if isLetter(r) {
		return 1
	}
	// Se la runa in input è un carattere underscore vai allo stato di non
	// accettazione q2
	if r == '_' || isDigit(r) {
		return 2
	}

	// Vai allo stato di non accettazione q2
	return 2
}

// gotoQ1 definisce le transizione dell'automa quando si trova nello stato q1
func gotoQ1(r rune) rune {
	// Se la runa in input è una lettera dell'alfabeto vai allo stato di
	// accettazione q3
	if isLetter(r) {
		return 3
	}
	// Se la runa in input è un carattere underscore rimani nello stato di
	// accettazione q1
	if r == '_' || isDigit(r) {
		return 1
	}

	// Vai allo stato di non accettazione q2
	return 2
}

// gotoQ2 definisce la transizione dell'automa quando si trova nello stato q2
func gotoQ2(r rune) rune {
	// Se la runa in input è una lettera dell'alfabeto rimani nello stato
	// di non accettazione q2
	if isLetter(r) {
		return 2
	}
	// Se la runa in input è un carattere underscore rimani nello stato di
	// non accettazione q2
	if r == '_' || isDigit(r) {
		return 2
	}

	// Per ogni altro simbolo dell'alfabeto rimani nello stato di non
	// accettazione q2
	return 2
}

// gotoQ3 definisce la transizione dell'automa quando si trova nello stato q3
func gotoQ3(r rune) rune {
	// Se la runa in input è una lettera dell'alfabeto rimani nello stato
	// di accettazione q3
	if isLetter(r) {
		return 3
	}
	// Se la runa in input è un carattere di underscore rimani nello stato
	// di accettazione q3
	if r == '_' || isDigit(r) {
		return 3
	}

	// Per ogni altro simbolo dell'alfabeto vai allo stato di non accettazione
	// q2
	return 2
}

// isLetter restituisce vero se la runa r è una lettera dell'alfabeto, sia maiuscola che minuscola.
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isDigit restituisce vero se la runa r è una cifra da 0 a 9.
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func main() {
	// Verifica se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: parserc <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando
	filename := os.Args[1]

	// Legge il contenuto del file in un buffer di byte
	buffer, err := readFileToBuffer(filename)
	if err != nil {
		wrappedErr := fmt.Errorf("error reading file: %w", err)
		fmt.Println(wrappedErr)
		os.Exit(1)
	}

	// Legge le parole dal buffer e avvia l'automa per determinare quali parole
	// sono accettate
	val, nval := stringsFromBuffer(buffer)
	fmt.Printf("Valide %d, invalide %d\n", val, nval)
}
