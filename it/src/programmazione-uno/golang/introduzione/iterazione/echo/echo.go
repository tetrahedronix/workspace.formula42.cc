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
/*
La maggior parte dei programmi lavora su dei dati di input per risolvere i
problemi. Questi dati possono provenire da diverse sorgenti: sono forniti dal
programma stesso; provengono da canali esterni (un file, una connessione di
rete, l'output di un altro programma, la tastiera dell'utente, gli argomenti
della riga di comando, ecc.).

Il pacchetto OS

Le funzionalità del sistema operativo sono disponibili al programmatore
tramite il pacchetto os. Si tratta di una interfaccia software indipendente
dalla macchina che rispetta lo standard Unix, seppure l'apparato disposto per
gestire gli errori è stato implementato osservando le caratteristiche del
linguaggio Go.

Per esempio in questo pacchetto si trova la variabile Args, una slice di
stringhe che conserva tutti gli argomenti presenti sulla riga di comando,
partendo dal nome del programma:

	var Args []string

Si usa la stringa os.Args ovunque si voglia utilizzare questa variabile al
di fuori del pacchetto os.

La slice

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
il programma stampi la prima stringa passata come argomento. Può essere d'aiuto
leggere il testo seguente, tratto dalla man page:

Il comando echo visualizza una linea di testo. La sintassi è la seguente

echo [SHORT-OPTION]...[STRING]...
echo LONG-OPTION

Fa l'eco di STRINGA(e) sullo standard output. Sono accettate le seguenti opzioni
corte:

-m

	non emette il ritorno a capo finale.

-e

	abilita l'interpretazione degli escape della barra rovesciata

-E

	disabilita l'interpretazione degli escape della barra rovesciata
	(comportamento di default)

--help

	visualizza questo messaggio di aiuto e termina

--version

	emette l'informazione sulla versione e termina

Se -e è attiva, sono riconosciute le seguenti sequenze:

\\		backslash

\a 		allerta sonora (BEL)

\b		backspace

\c		non produce ulteriore output

\e		escape

\f		avanzamento modulo

\n		nuova linea

\r		ritorno a capo

\t		tabulazione orizzontale

\v		tabulazione verticale

\0NNN	byte con valore ottale NNN (da 1 a 3 cifre)

\xHH	byte con valore esadecimale HH (da 1 a 2 cifre)

NOTA:	la tua shell può avere la sua versione di echo che disolito prevale
sulla versione descritta qui. Fare riferimento alla documentazione della shell
per i dettagli sulle opzioni supportate.

NOTA:	la printf(1) è un'alternativa preferita, non presenta problemi
nell'emettere stringhe simili a opzioni.
*/

/*
echo emula la funzione principale del comando echo: stampa tutte le stringhe
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
Il costrutto for viene usato per fare ogni tipo di iteraione, quindi l'istruzione
for ha diverse forme nel linguaggio, per esempio la più comune è la seguente

	for inizializzazione; condizione; post {

	}

Non sono necessarie le parentesi tonde attorno alle tre componenti di controllo.
Sono invece perentorie le parentesi graffe del blocco: la prima deve comparire
sulla stessa riga dell'istruzione for, la seconda chiude il blocco su una linea
a se stante. È un'istruzione for privata del suo blocco di codice, anche vuoto,
è illegale:

	for i:=0; i<10; i++;

È legale un'istruzione for con un blocco di codice vuoto:

	for i:=0; i<10; i++ {

	}

Le componenti inizializzazione, condizione, post sono opzionali. L'inizializione
è un comando eseguito una sola volta prima che il ciclo inizi. Può essere una
dichiarazione e definizione di variabile in linea (una tuple, non è ammessa la
parola chiave var). In tal caso, la variabile è locale e pertanto ha precedenza
su un altro identificatore con lo stesso nome esterno al blocco for. Tuttavia
omettendo l'inizializzazione è possibile impostare e utilizzare un
identificatore esterno per condizionare il ciclo e nella operazione finale
(post), come nel seguente breve codice d'esempio

	var i int = 1

	for ; i < 10; i++ {

	}

L'inizializzazione può altresì essere un incremento oppure una chiamata di
funzione

	for A(); ; {

	}
	...

	func A() {

	}

La condizione è un'espressione booleana, cioè significa che assume una logica a
due valori, vero oppure falso. Viene valutata all'inizio del ciclo e ad ogni
iterazione: se il valore è vero, i comandi presenti nel blocco di codice del
ciclo sono eseguiti.

Il comando di finalizzazione (post) è eseguito per ultimo dopo ogni altra
istruzione presente nel ciclo. In generale si tratta di un'espressione (oppure
un comando di incremento/decremento) che cambia il valore alla variabile di
controllo definita in precedenza. Dopo il post il programma ritorna alla
espressione booleana booleana per stabilire se ci sono le condizioni per
eseguire un'altra iterazione. Il ciclo quindi termina se tale espressione
assume il valore falso.

Ciclo infinito

Un ciclo è infinito se viene omessa la condizione, a meno che all'interno del
blocco dell'istruzione for compaiano istruzioni che causano un'interruzione,
per esempio le istruzioni break, return oppure la funzione os.Exit()


*/
