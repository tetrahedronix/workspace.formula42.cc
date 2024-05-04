/**
Creative Commons Legal Code

CC0 1.0 Universal

    CREATIVE COMMONS CORPORATION IS NOT A LAW FIRM AND DOES NOT PROVIDE
    LEGAL SERVICES. DISTRIBUTION OF THIS DOCUMENT DOES NOT CREATE AN
    ATTORNEY-CLIENT RELATIONSHIP. CREATIVE COMMONS PROVIDES THIS
    INFORMATION ON AN "AS-IS" BASIS. CREATIVE COMMONS MAKES NO WARRANTIES
    REGARDING THE USE OF THIS DOCUMENT OR THE INFORMATION OR WORKS
    PROVIDED HEREUNDER, AND DISCLAIMS LIABILITY FOR DAMAGES RESULTING FROM
    THE USE OF THIS DOCUMENT OR THE INFORMATION OR WORKS PROVIDED
    HEREUNDER.
*/
// E10.2.6, pag 536
// Da "Foundations of Computer Science: C Edition", Aho, Ullmann
//
// Progettare un automa che dice se una data stringa di caratteri è un pronome
// "he", "his", "him", "she", "her" oppure "hers"
package main

import (
	"bufio"
	"bytes"
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

// stringsFromBuffer legge le parole da un buffer di byte e chiama la funzione
// searchPronom per determinare se ogni parola è un pronome inglese: "he",
// "his", "him", "she", "her" oppure "hers"
func stringsFromBuffer(buffer []byte) {
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

		// Chiama la funzione searchPronom per determinare se la parola
		// è un pronome inglese.
		if searchPronom(word) {
			fmt.Println(word, "OK")
		} else {
			fmt.Println(word, "NO")
		}

	}
}

// Tabella di transizione senza riconoscimento di spazio bianco
//
//     | h  | e  |  r | i  | s  | m  |  S\{h,e,r,i,s,m}
// ----|----|----|----|----|----|----|------------------
// +q0 | q1 | q0 | q0 | q0 | q6 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
//  q1 | q0 | q2 | q0 | q5 | q0 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
// *q2 | q0 | q0 | q3 | q0 | q0 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
// *q3 | q0 | q0 | q0 | q0 | q4 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
// *q4 | q0 | q0 | q0 | q0 | q0 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
//  q5 | q0 | q0 | q0 | q0 | q4 | q7 |    q0
// ----|----|----|----|----|----|----|------------------
//  q6 | q1 | q0 | q0 | q0 | q0 | q0 |    q0
// ----|----|----|----|----|----|----|------------------
//  q7 | q0 | q0 | q0 | q0 | q0 | q0 |    q0
//  ----------------------------------------------------

func searchPronom(w string) bool {
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
		// Passa dallo stato q4 allo stato q0
		case 4:
			q = 0
		// Passa dallo stato q5 allo stato q restituito da gotoQ5
		case 5:
			q = gotoQ5(r)
		// Passa dallo stato q6 allo stato q restituito da gotoQ6
		case 6:
			q = gotoQ6(r)
		// Passa dallo stato q7 allo stato q0
		case 7:
			q = 0
		}
	}

	return q == 2 || q == 3 || q == 4 || q == 7
}

// gotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
func gotoQ0(r rune) rune {

	if r == 'h' {
		return 1
	}

	if r == 's' {
		return 6
	}

	return 0
}

// gotoQ1 definisce le transizioni dell'automa quando si trova nello stato q1
func gotoQ1(r rune) rune {

	if r == 'e' {
		return 2
	}

	if r == 'i' {
		return 5
	}

	return 0
}

// gotoQ2 definisce le transizioni dell'automa quando si trova nello stato q2
func gotoQ2(r rune) rune {

	if r == 'r' {
		return 3
	}

	return 0
}

// gotoQ3 definisce le transizioni dell'automa quando si trova nello stato q3
func gotoQ3(r rune) rune {

	if r == 's' {
		return 4
	}

	return 0
}

// gotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
// (da implementare per l'automa che esce con il carattere di spazio bianco)
// func gotoQ4(r rune) rune {

// 	return 0
// }

// gotoQ5 definisce le transizioni dell'automa quando si trova nello stato q5
func gotoQ5(r rune) rune {

	if r == 's' {
		return 4
	}

	if r == 'm' {
		return 7
	}

	return 0
}

// gotoQ6 definisce le transizioni dell'automa quando si trova nello stato q6
func gotoQ6(r rune) rune {

	if r == 'h' {
		return 1
	}

	return 0
}

// gotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
// (da implementare per l'automa che esce con il carattere di spazio bianco)
// func gotoQ7(r rune) rune {

// 	return 0
// }

func main() {

	// Verifica se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: pronomi <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando
	filename := os.Args[1]

	// Legge il contenuto del file in un buffer di byte
	buffer, err := readFileToBuffer(filename)
	if err != nil {
		wrapperErr := fmt.Errorf("error reading file: %w", err)
		fmt.Println(wrapperErr)
	}

	// Legge le parole dal buffer e avvia l'automa per determinare quali
	// parole sono accettate
	stringsFromBuffer(buffer)
}
