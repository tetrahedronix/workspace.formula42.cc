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
// Da "Foundations of Computer Science: C Edition", Aho, Ullmann
//
// # Progettare un automa per leggere stringhe di 0 e 1, e
//
// a) esegua il controllo se la sequenza letta finora ha parità pari (cioè
// ci sono stati un numero pari di 1). In particolare, l'automa accetta se
// la stringha finora letta ha parità pari e rifiuta se ha parità dispari.
//
// Controllare che non ci siano più di due 1 consecutivi. In altre parole,
// accettare a meno che 111 non sia una sottostringa della stringa di input
// finora letta
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// readFIleToBuffer legge il contenuto di un file dato il suo nome
// e restituisce il contenuto come un buffer di byte.
func readFIleToBuffer(f string) ([]byte, error) {
	// Apre il file in lettura
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Legge il contenuto del file in un buffer
	buffer, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// stringsFromBuffer legge le parole da un buffer di byte e chiama
// la funzione searchPatternA per determinare se ogni parola presenta
// un numero pari di 1 e searchPatternB per determinare se ogni parola
// presenta al più due simboli 1 consecutivi.
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

		// Chiama la funzione searchPatternA per determire se la parola ha
		// un numero pari di 1.
		if searchPatternA(word) {
			fmt.Print(word, " OK")
		} else {
			fmt.Print(word, " NO")
		}

		if searchPatternB(word) {
			fmt.Println("", "OK")
		} else {
			fmt.Println("", "NO")
		}

	}
}

// searchPatternA implementa un automa che accetta le parole che
// contengono un numero pari di 1.
//
// L'automa ha la seguente tabella di transizione
//
//	       |   0   |   1
//	=======|=======|=======
//	+*q0   |  q0   |  q1
//	-----------------------
//	  q1   |  q1   |  q0
//	-----------------------
func searchPatternA(w string) bool {

	var q rune // Stato corrente dell'automa

	for _, r := range w {
		// Stato q0
		if q == 0 {
			q = gotoQ0(r)
			// Stato q1
		} else if q == 1 {
			q = gotoQ1(r)
		}
	}

	// L'automa accetta se si trova nello stato finale q0
	return q == 0
}

func searchPatternB(w string) bool {

	var q rune // Stato corrente dell'automa

	for _, r := range w {

		switch q {
		// Passa dallo stato q0	a q
		case 0:
			q = gotoQ0(r)
		// Passa dallo stato q1 a q
		case 1:
			q = gotoQ1b(r)
		// Passa dallo stato q2 a q
		case 2:
			q = gotoQ2(r)
		// Permane nello stato q3
		case 3:
			q = 3
		}
		// if q == 0 {
		// 	q = gotoQ0(r)
		// } else if q == 1 {
		// 	q = gotoQ1b(r)
		// } else if q == 2 {
		// 	q = gotoQ2(r)
		// } else if q == 3 {
		// 	q = gotoQ3(r)
		// }

	}

	// L'automa B accetta se si trova in uno degli stati accettati (0, 1, 2).
	// Utilizziamo q != 3 per verificare se lo stato corrente è diverso da 3,
	// che rappresenta (q=3) uno stato di rifiuto.
	return q != 3 // Equivalente a return q == 0 || q == 1 || q == 2

}

// gotoQ0 definisce le transizioni dell'automa A (e B) nello stato q0.
func gotoQ0(r rune) rune {

	if r == '0' {
		// Rimane nello stato q0
		return 0
	}
	// Passa allo stato q1
	return 1
}

// gotoQ1 definisce le transizioni dell'automa A nello stato q1.
func gotoQ1(r rune) rune {

	if r == '0' {
		// Rimane nello stato q1
		return 1
	}
	// Torna allo stato q0
	return 0
}

// gotoQ1 definisce le transizioni dell'automa B nello stato q1.
func gotoQ1b(r rune) rune {

	if r == '0' {
		// vai allo stato q0
		return 0
	}

	return 2
}

func gotoQ2(r rune) rune {

	if r == '0' {
		return 0
	}

	return 3
}

// func gotoQ3(r rune) rune {

// 	return 3

// }

func main() {
	// Verifica se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run unopari <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando
	filename := os.Args[1]

	// Legge il contenuto del file in un buffer di byte
	buffer, err := readFIleToBuffer(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Legge le parole dal buffer e avvia l'automa per determinare quali
	// parole sono accettate.
	stringsFromBuffer(buffer)
}
