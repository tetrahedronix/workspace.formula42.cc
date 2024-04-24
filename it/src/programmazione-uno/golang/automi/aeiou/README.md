## Istruzioni

1. Il seguente comando compila il programma `aeiou` con il compilatore `go` e
produce l'eseguibile `aeiou`

```
go build aeiou.go

```

oppure per eseguire il programma `aeiou` senza costruire l'eseguibile:

```
go run aeiou.go

```

2. Per verificare la correttezza del programma, si consiglia di eseguire il
*testing* unitario utilizzando il seguente comando:

```
go test

```

## Definizione dell'automa

L'automa `aeiou` è una quintupla

```
A=(Q, Σ, 𝛔, q0, F)

```

Dove

1. Q={q0,q1,q2,q3,q4,q5,q6} è l'insieme degli stati dell'automa. 
2. Σ rappresenta l'alfabeto latino.
3. δ è la funzione di transizione.
4. q0 indica lo stato iniziale.
5. F={q5} è l'insieme degli stati finali

## Descrizione dell'automa

L'automa è costituito da 7 stati, più uno stato speciale per rifiutare la parola
vuota (non previsto dalla definizione di automa deterministico). Lo stato
iniziale è `q0`. Ci sono cinque stati `q1`, `q2`, `q3`, `q4`, `q5` dove l'automa
transita o permane se viene ricevuta in input rispettivamente la vocale `a`, `e`,
`i`, `o`, `u`. Vi è uno stato speciale `q6` di non accettazione, dove l'automa
transita ogni qualvolta viene ricevuta in input una vocale che è di ordine
lessicografico superiore alla vocale rappresentata dal corrente stato dell'automa.
Infine, lo stato di accettazione `q5` permette all'automa di terminare il
riconoscimento della parola con le vocali in ordine lessicografico.

La funzione di transizione δ accetta come input una coppia `(q,σ)` dove q è
lo stato attuale dell'automa e σ è un simbolo d'input dell'alfabeto Σ. Ci sono
regole specifiche che determinano il comportamento dell'automa in base al 
valore restituito dalla funzione δ, per esempio `δ(q0, a)` restituisce `q1`,
`δ(q5, u)` restituisce `q5`, e così via (vd. tabella di transizione).

Gli stati `q1`,`q2`,`q3`,`q4` e `q5` rappresentano le situazioni in cui l'automa
si trova quando viene incontrata rispettivamente la vocale "a", "e", "i", "o",
"u". La funzione di transizione δ deve quindi mappare questi stati dell'automa
alle vocali e restituire uno stato di avanzamento in base all'ordine
lessicografico. In caso contrario, la funzione δ permette all'automa di 
transitare allo stato speciale `q6` di non accettazione. Esistono diversi
casi possibili di transizione in `q6`, per esempio qualora venga valutata una
`δ(q4,a)`: nello stato `q4` su input `a` l'automa ha già incontrato le prime
quattro vocali e non è ammessa una ripetizione di questo simbolo.

## Comportamento atteso del programma

Il programma `aeiou` implementa un automa a stati finiti per verificare se
una parola contiene tutte e cinque le vocali ('a', 'e', 'i', 'o', 'u') in
ordine lessicografico. Ecco un riassunto del comportamento atteso del programma:

1. **Inizializzazione**: Il programma legge il contenuto di un file specificato
come argomento della riga di comando e lo converte in un buffer di byte.

2. **Analisi delle parole**: Le parole vengono estratte dal buffer e vengono
testate dall'automa implementato per determinare se contengono le vocali in
ordine.

3. **Automazione a stati**: L'automa ha uno stato iniziale e transizioni
definite per ogni carattere delle vocali. Le transizioni sono progettate per
accettare solo parole che contengono le vocali in ordine lessicografico.

4. **Verifica delle vocali**: L'automa inizia cercando la 'a' come prima
vocale, poi procede con 'e', 'i', 'o' e infine 'u'. Se tutte le vocali sono
trovate in ordine, la parola viene accettata.

5. **Restituzione del risultato**: Il programma determina se ogni parola nel
file soddisfa i criteri di accettazione dell'automa e stampa il risultato.

In sostanza, il programma serve a verificare se le parole contenute in un file
rispettano un certo schema di presenza delle vocali in ordine lessicografico.

## Diagramma di transizione

![Diagramma di transizione automa `aeiou`](./asset/diagramma.svg)

## Tabella di transizione

```
       |   a   |   e   |   i   |   o   |   u   |  Σ\{a,e,i,o,u}
=======|=======|=======|=======|=======|=======|===========
 +q0   |  q1   |  q6   |  q6   |  q6   |  q6   |    q0
-----------------------------------------------------------
  q1   |  q1   |  q2   |  q6   |  q6   |  q6   |    q1
-----------------------------------------------------------
  q2   |  q6   |  q2   |  q3   |  q6   |  q6   |    q2
-----------------------------------------------------------
  q3   |  q6   |  q6   |  q3   |  q4   |  q6   |    q3
-----------------------------------------------------------
  q4   |  q6   |  q6   |  q6   |  q4   |  q5   |    q4
-----------------------------------------------------------
 *q5   |  q5   |  q5   |  q5   |  q5   |  q5   |    q5
-----------------------------------------------------------
  q6   |  q6   |  q6   |  q6   |  q6   |  q6   |    q6
```

## Esempi

Esempi di parole accettate dall'automa presenti in `aeiou_test.go`

1. `""`. La parola vuota ε non è accettata dall'automa. In una implementazione
di un automa deterministico per riconoscere le vocali in ordine lessicografico,
non è prevista l'accettazione della parola vuota ε. Questo è dovuto al fatto che
un automa deterministico opera su una serie di stati e transizioni definite,
senza ε-transizioni che consentono il passaggio diretto tra gli stati in assenza
di input. Poiché la parola vuota non contiene alcun input, non esiste uno stato
di accettazione che possa riconoscerla. Di conseguenza, l'automa `aeiou` non
accetta tale parola e opera esclusivamente su stringhe non vuote contenenti le
vocali in ordine lessicografico. Tuttavia, nell'implementazione di `aeiou`
viene aggiunto uno *stato fittizio* che rappresenta una situazione di
non-accettazione raggiungibile dall'inizio solo se l'input è la parola vuota ε.
Questa soluzione non viola la definizione di automa a stati finiti deterministico,
poiché quando viene controllata la presenza della parola vuota, l'automa non è
stato ancora "inizializzato".

2. "abstemious". Questa parola contiene tutte e cinque le vocali in ordine
lessicografico e pertanto è accettata dall'automa nel seguente modo:
  * Su input `a`, avviene la transizione `δ(q0,a)=q1`, quindi l'automa passa
  dallo stato `q0` allo stato `q1`.
  * Quando vengono lette le consonanti `b`, `s`, `t`, l'automa rimane nello
  stato `q1` perché le consonanti non influenzano l'ordine lessicografico
  richiesto per le vocali: pertanto si ha `δ(q1,b)=q1`, `δ(q1,s)=q1`, `δ(q1,t)=q1`.
  * Su input `e`, avviene la transizione `δ(q1,e)=q2`. Nello stato `q2`,
  l'ordine lessicografico è rispettato, anche se la sequenza di vocali cercata
  dall'automa è parziale.
  * Viene letto il simbolo `m`, con la transizione `δ(q2,m)=q2`.
  * Su input `i`, l'automa si sposta allo stato `q3` con la transizione
  `δ(q2,i)=q3`. In `q3`, l'automa presumibilmente si aspetta di poter leggere
  la vocale successiva `o` dopo aver incontrato le prime tre vocali.
  * I due simboli successivi sono le vocali `o` e `u`, che determinano le
  transizioni `δ(q3,o)=q4` e `δ(q4,u)=q5`. Ora l'automa si trova nello stato
  finale di accettazione. Non è necessario leggere il simbolo successivo dato
  che è stata rilevata una sequenza completa di vocali in ordine lessicografico.

3. "zoo". Questa parola è rifiutata dall'automa in quanto non viene mai
raggiunto lo stato finale `q5`. Su input `z` avviene la transizione `δ(q0,z)=q0`.
Successivamente, la transizione `δ(q0,o)=q6` su input `o` porta l'automa
allo stato q6 di non accettazione, dal quale non può più uscire.

## Codice sorgente

```Go
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
	// Apre il file in lettura.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Legge il contenuto del file in un buffer.
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
	// Creazione di un nuovo reader basato sul buffer.
	reader := bytes.NewReader(buffer)

	// Creazione di un nuovo scanner che leggerà dal reader.
	scanner := bufio.NewScanner(reader)

	// Impostazione dello scanner per scansionare le parole.
	scanner.Split(bufio.ScanWords)

	// Ciclo finché ci sono parole da leggere.
	for scanner.Scan() {
		// Legge la parola dallo scanner.
		word := scanner.Text()

		// Chiama la funzione searchPattern per determinare se la parola ha le
		// vocali in ordine lessicografico.
		searchPattern(word)
		//testWord(word)
	}

	// Controlla se ci sono errori durante la scansione.
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
//			  q2   |  q6   |  q2   |  q3   |  q6   |  q6   |    q2
//			-----------------------------------------------------------
//			  q3   |  q6   |  q6   |  q3   |  q4   |  q6   |    q3
//			-----------------------------------------------------------
//			  q4   |  q6   |  q6   |  q6   |  q4   |  q5   |    q4
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
		return false // Se la parola è vuota, restituisce true.
	}

	// INizializzazione dell'automa.
	var q rune

	for _, r := range w {

		switch q {
		case 0:
			// Passa dallo stato q0 allo stato q su input r.
			q = gotoQ0(r)
		case 1:
			// Passa dallo stato q1 allo stato q su input r.
			q = gotoQ1(r)
		case 2:
			// Passa dallo stato q2 allo stato q su input r.
			q = gotoQ2(r)
		case 3:
			// Passa dallo stato q3 allo stato q su input r.
			q = gotoQ3(r)
		case 4:
			// Passa dallo stato q4 allo stato q su input r.
			q = gotoQ4(r)
		case 5:
			// Rimane nello stato di accettazione 5.
			q = 5
		}

		// Se l'automa si trova nello stato q6 di non accettazione, restituire
		// falso (la parola è rifiutata). Questo controllo è necessario per
		// interrompere l'analisi.
		if q == 6 {
			return false
		}
	}

	// Se l'automa termina l'esecuzione nello stato di accettazione 5,
	// restituire vero (la parola è accettata).
	return q == 5
}

// qotoQ0 definisce le transizioni dell'automa quando si trova nello stato q0
func gotoQ0(r rune) rune {

	if r == 'a' {
		// Transita allo stato q1.
		return 1
	}

	if r == 'e' || r == 'i' || r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione.
		return 6
	}

	// Per ogni altro carattere mantiene lo stato q0.
	return 0
}

// gotoQ1 definisce le transizione dell'automa quando si trova nello stato q1.
func gotoQ1(r rune) rune {

	if r == 'e' {
		// Transita allo stato q2.
		return 2
	}

	if r == 'i' || r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione.
		return 6
	}

	// Per ogni altro carattere mantiene lo stato q1
	return 1
}

// gotoQ2 definisce la transizione dell'automa quando si trova nello stato q2.
func gotoQ2(r rune) rune {

	if r == 'i' {
		// Transita allo stato q3.
		return 3
	}

	if r == 'a' || r == 'o' || r == 'u' {
		// Transita allo stato q6 di non accettazione.
		return 6
	}

	// Mantieni lo stato q2.
	return 2
}

// gotoQ3 definisce la transizione dell'automa quando si trova nello stato q3.
func gotoQ3(r rune) rune {

	if r == 'o' {
		// Transita allo stato q4.
		return 4
	}

	if r == 'a' || r == 'e' || r == 'u' {
		// Transita allo stato q6 di non accettazione.
		return 6
	}

	// Mantieni lo stato q3.
	return 3
}

// gotoQ4 definisce la transizione dell'automa quando si trova nello stato q4.
func gotoQ4(r rune) rune {

	if r == 'u' {
		// Transita allo stato di accettazione q5.
		return 5
	}

	if r == 'a' || r == 'e' || r == 'i' {
		// Transita allo stato q6 di non accettazione
		return 6
	}

	// Mantieni lo stato q4.
	return 4
}

// findChar cerca la prima occorrenza del carattere 'r' nella stringa 's'.
// Restituisce l'indice della prima occorrenza e true se il carattere è stato trovato,
// altrimenti restituisce 0 e false.
func findChar(s string, r rune) (i int, b bool) {
	// Cicla attraverso la stringa 's'
	for i, c := range s {
		// Controlla se il carattere corrente è uguale a 'r'.
		if c == r {
			// Se è stato trovato il carattere, restituisce l'indice e true.
			return i, true
		}
	}
	// Se il carattere non è stato trovato, restituisce 0 e false.
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
		// Cerca la posizione della lettera 'e' nella parte della parola dopo la lettera 'a'.
		if i, ok := findChar(w[i:], 'e'); ok {
			// Cerca la posizione della lettera 'i' nella parte della parola dopo la lettera 'e'.
			if i, ok := findChar(w[i:], 'i'); ok {
				// Cerca la posizione della lettera 'o' nella parte della parola dopo la lettera 'i'.
				if i, ok := findChar(w[i:], 'o'); ok {
					// Cerca la posizione della lettera 'u' nella parte della parola dopo la lettera 'o'.
					if _, ok := findChar(w[i:], 'u'); ok {
						// Se tutte le vocali sono state trovate in ordine, restituisce true.
						return true
					}
				}
			}
		}
	}
	// Se la parola non contiene tutte e cinque le vocali in ordine,
	// restituisce false.
	return false
}

func main() {
	// Verifica se è stato fornito il nome del file come argomento.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	// Ottiene il nome del file dall'argomento della riga di comando.
	filename := os.Args[1]

	// Legge il contenuto del file in un buffer di byte.
	buffer, err := readFileToBuffer(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Legge le parole dal buffer e applica l'automa per determinare quali
	// parole sono accettate.
	stringsFromBuffer(buffer)
}
```