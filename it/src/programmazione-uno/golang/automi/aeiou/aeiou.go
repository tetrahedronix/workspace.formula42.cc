// E10.2, pag. 531
// Da "Foundations of Computer Science: C Edition", Aho, Ullmann
//
// Scrivere un semplice programma per esaminare una stringa di caratteri e
// stabilire se contiene tutte e cinque le vocali in ordine.
// Partendo dall'inizio della stringa, il programma cerca prima la 'a'.
// Si può dire che il programma è nello "stato 0" finché non trova una 'a',
// dopodiché passa allo "stato 1". Nello stato 1, cerca la 'e' e, una volta
// trovata, passa allo "stato 2". Continua in questo modo fino a raggiungere
// lo "stato 4", in cui cerca la 'u'. Se trova la 'u', significa che la
// parola contiene tutte e cinque le vocali in ordine, e il programma può
// passare allo "stato di accettazione 5". Non è necessario analizzare il
// resto della parola, poiché il programma sa già che è valida,
// indipendentemente da ciò che segue la 'u'.
//
// Note per il Copyright:
//   - l'automa searchPattern è rilasciato sotto GNU GPL.
//   - l'implentazione dell'automa findChar è una traduzione in Golang
//     dell'automa findChar dell'opera "Foundations of Computer Science"
//     di Aho e Ullmann.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// La funzione readFileToBuffer legge il contenuto di un file dato il suo nome
// e restituisce il contenuto come un buffer di byte.
func readFileToBuffer(filename string) ([]byte, error) {
	// Apre il file in lettura
	file, err := os.Open(filename)
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

// La funzione stringsFromBuffer legge le parole da un buffer di byte e chiama
// la funzione searchPattern per determinare se ogni parola ha le vocali in
// ordine lessicografico.
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

		// Chiama la funzione searchPattern per determinare se la parola ha le
		// vocali in ordine lessicografico
		searchPattern(word)
		//testWord(word)
	}

	// Controlla se ci sono errori durante la scansione
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading buffer: %v\n", err)
	}
}

// La funzione searchPattern implementa un automa che accetta le parole che
// hanno le vocali in ordine lessicografico.
// L'automa ha la seguente tabella di transizione
//
//			       |   a   |   e   |   i   |   o   |   u   |  Σ\{a,e,i,o,u}
//			=======|=======|=======|=======|=======|=======|===========
//			 +q0   |  q1   |  q6   |  q6   |  q6   |  q6   |    q0
//			-----------------------------------------------------------
//			  q1   |  q1   |  q2   |  q6   |  q6   |  q6   |    q1
//			-----------------------------------------------------------
//			  q2   |  q2   |  q2   |  q3   |  q6   |  q6   |    q2
//			-----------------------------------------------------------
//			  q3   |  q3   |  q3   |  q3   |  q4   |  q6   |    q3
//			-----------------------------------------------------------
//			  q4   |  q4   |  q4   |  q4   |  q4   |  q5   |    q4
//			-----------------------------------------------------------
//			 *q5   |  q5   |  q5   |  q5   |  q5   |  q5   |    q5
//		  -----------------------------------------------------------
//		      q6   |  q6   |  q6   |  q6   |  q6   |  q6   |    q6
//
//	 q0: stato iniziale
//	 q5: stato di accettazione
//	 q6: stato speciale di non accettazione
//
// Rispetto all'automa sviluppato nel libro di Aho & C., l'automa searchPattern
// accetta parole solo se presentano le vocali in ordine lessicografico. Quindi
// la parola "sacrilegious" non è accettata, mentre la parola "aaeeiioouu" è
// valida (vd. aeiou_test.go per ulteriori esempi).
func searchPattern(w string) bool {

	if len(w) == 0 {
		return true // Se la parola è vuota, restituisce true
	}

	var q rune

	for _, r := range w {
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
			q = 5
		}

		if q == 6 {
			return false
		}
	}

	return q == 5
}

// qotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
func gotoQ0(r rune) rune {

	if r == 'a' {
		// Transita allo stato q1
		return 1
	}

	if r == 'e' || r == 'i' || r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione
		return 6
	}

	// Mantieni lo stato q0
	return 0
}

// gotoQ1 definisce le transizione dell'automa quando si trova nello stato q1
func gotoQ1(r rune) rune {

	if r == 'e' {
		// Transita allo stato q2
		return 2
	}

	if r == 'i' || r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione
		return 6
	}

	// Mantieni lo stato q1
	return 1
}

// gotoQ2 definisce la transizione dell'automa quando si trova nello stato q2
func gotoQ2(r rune) rune {

	if r == 'i' {
		// Transita allo stato q3
		return 3
	}

	if r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione
		return 6
	}

	// Mantieni lo stato q2
	return 2
}

// gotoQ3 definisce la transizione dell'automa quando si trova nello stato q3
func gotoQ3(r rune) rune {

	if r == 'o' {
		// Transita allo stato q4
		return 4
	}

	if r == 'u' {
		// Transita allo stato q6 di non accettazione
		return 6
	}

	// Mantieni lo stato q3
	return 3
}

// gotoQ4 definisce la transizione dell'automa quando si trova nello stato q4
func gotoQ4(r rune) rune {

	if r == 'u' {
		// Transita allo stato di accettazione q5
		return 5
	}

	// Mantieni lo stato q4
	return 4
}

// findChar cerca la prima occorrenza del carattere 'r' nella stringa 's'.
// Restituisce l'indice della prima occorrenza e true se il carattere è stato trovato,
// altrimenti restituisce 0 e false.
func findChar(s string, r rune) (i int, b bool) {
	// Cicla attraverso la stringa 's'
	for i, c := range s {
		// Controlla se il carattere corrente è uguale a 'r'
		if c == r {
			// Se è stato trovato il carattere, restituisce l'indice e true
			return i, true
		}
	}
	// Se il carattere non è stato trovato, restituisce 0 e false
	return i, false
}

// testWord verifica se una parola 'w' contiene una sottosequenza composta da
// tutte e cinque le vocali ('a', 'e', 'i', 'o', 'u') in ordine lessicografico.
// Utilizza la funzione findChar per cercare le vocali nella parola.
// Restituisce true se la parola contiene tutte e cinque le vocali in ordine,
// altrimenti restituisce false.
//
// Questo automa è basato sulla funzione testWord scritta in linguaggio C
// di Fig. 10.2, pag. 532, cap. 10, "Foundations of Computer Science", Aho,
// Ullmann.
//
// L'automa implementato nel libro accetta la parola "sacrilegious", 'aieiou',
// ma non la parola vuota "". Pertanto il test fallisce se vengono passate
// come input in `aeiou_test.go`
func testWord(w string) bool {
	// Cerca la posizione della lettera 'a' nella parola 'w'
	if i, ok := findChar(w[0:], 'a'); ok {
		// Cerca la posizione della lettera 'e' nella parte della parola dopo la lettera 'a'
		if i, ok := findChar(w[i:], 'e'); ok {
			// Cerca la posizione della lettera 'i' nella parte della parola dopo la lettera 'e'
			if i, ok := findChar(w[i:], 'i'); ok {
				// Cerca la posizione della lettera 'o' nella parte della parola dopo la lettera 'i'
				if i, ok := findChar(w[i:], 'o'); ok {
					// Cerca la posizione della lettera 'u' nella parte della parola dopo la lettera 'o'
					if _, ok := findChar(w[i:], 'u'); ok {
						// Se tutte le vocali sono state trovate in ordine, restituisce true
						return true
					}
				}
			}
		}
	}
	// Se la parola non contiene tutte e cinque le vocali in ordine, restituisce false
	return false
}

func main() {
	// Verifica se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando
	filename := os.Args[1]

	// Legge il contenuto del file in un buffer di byte
	buffer, err := readFileToBuffer(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Legge le parole dal buffer e applica l'automa per determinare quali
	// parole sono accettate.
	stringsFromBuffer(buffer)
}
