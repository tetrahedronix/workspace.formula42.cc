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
/*
La maggior parte dei programmi lavora su dei dati di input per risolvere i
problemi. Questi dati possono provenire da diverse sorgenti: sono forniti dal
programma stesso; provengono da canali esterni (un file, una connessione di
rete, l'output di un altro programma, la tastiera dell'utente, gli argomenti
della riga di comando, ecc.).

## Il pacchetto OS

Le funzionalità del sistema operativo sono disponibili al programmatore
tramite il pacchetto `os`. Si tratta di una interfaccia software indipendente
dalla macchina che rispetta lo standard Unix, ma l'apparato disposto per
gestire gli errori è stato implementato osservando le caratteristiche del
linguaggio Go.

Per esempio in questo pacchetto si trova la variabile Args, una slice di
stringhe che conserva tutti gli argomenti presenti sulla riga di comando,
partendo dal nome del programma:

```
	var Args []string
```

Si usa la stringa os.Args ovunque si voglia utilizzare questa variabile al
di fuori del pacchetto os.

## La slice

Una slice è una sequenza di elementi la cui dimensione può variare durante
l'esecuzione del programma. Se s è una slice, allora ogni elemento è
raggiungibile per indirizzo con la scrittura s[i], dove i è un numero intero
maggiore o uguale a 0, mentre una porzione di sequenza è selezionata
specificando due numeri interi, per esempio s[m:n], che rispettivamente indicano
il primo elemento e l'ultimo (escluso) della porzione di slice.

Le seguenti istruzioni sono molto utili quando si lavora con le slice

	os.Args[0] è la variabile che contiene il nome del programma

	os.Args[1:len(os.Args)] è la lista degli argomenti specificati sulla riga
	dei comandi

	os.Args[1:] è la lista degli argomenti specificati sulla riga dei comandi,
	escluso il nome del programma, cioè l'elemento di indice 0 nella slice.

	os.Args[:] è la lista completa, cioè è equivalente a os.Args[0:len(os.Args)]
*/

/*
Implementare il comando Unix echo che stampa una linea di testo passata come
argomento. Non si richiede un'implementazione completa, è sufficiente che
il programma stampi la prima stringa passata come argomento.
*/

/*
efficient-echo emula la funzione principale del comando echo: stampa tutte le stringhe
passate sulla riga di comando, escluso il nome del programma stesso. Poi stampa
un carattere di ritorno a capo.
*/
package main

import (
	"fmt"
	"os"
)

// main ogni argomento passato al programma viene stampato in un ciclo, uno alla
// volta.
func main() {

	// Cicla all'interno della slice degli argomenti, partendo dal secondo
	// elemento, quello di indice 1, tralasciando così il nome del programma.
	for _, arg := range os.Args[1:] {
		// Stampa l'argomento arg seguito da uno spazio bianco:
		// L'operatore + applicato alle stringhe produce la loro concatenzione.
		fmt.Print(arg + " ") // Stesso come Print(arg, " ").
	}
	// Stampa un carattere speciale di ritorno a capo.
	fmt.Println()

}

/*
## Il ciclo for

In Go il costrutto `for` rappresenta l'unica istruzione di ciclo capace di coprire
le funzioni di `for` e `while` del linguaggio C. Non esiste invece l'equivalente
del `do-while`. La sua versatilità permette di esprimere ogni tipo di iterazione
attraverso tre varianti principali:

- **Forma completa**: for INIZIALIZZAZIONE; CONDIZIONE; POST { }
- **Forma condizionale**: for CONDIZIONE { }
- **Forma ifninita**: for { }

In tutte le varianti, le parentesi graffe che delimitano il blocco sono sempre
obbligatorie.

### Aspetti sintattici e strutturali

Nella forma completa, i tre componenti di controllo (inizializzazione, condizione
e post) non possono essere racchiusi tra parentesi tonde, poiché la grammatica del
linguaggio non lo consente. Soli in questa forma sono ammessi i punti e virgola
che separano le parti.

Un ciclo infinito può essere scritto come

```
for ; ; {
}
```

ma il compilatore lo formatterà automaticamente in

```
for {
}
```

Un ciclo è infinito se viene omessa la condizione, a meno che all'interno del
blocco dell'istruzione for compaiano istruzioni che causano un'interruzione,
per esempio le istruzioni break, return oppure la funzione os.Exit()

La forma più comune di utilizzo del for è quella con i tre componenti ben esplicitati:

```
for i := 0; i < 10; i++ {
    // corpo del ciclo
}
```

Un for privo di blocco di codice, anche se terminante con punto e virgola, non
è valido:

```
for i := 0; i < 10; i++; // errore
```

Viceversa, è lecito un ciclo con blocco vuoto:

```
for i := 0; i < 10; i++ {
}
```

*/
