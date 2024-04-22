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
 * Il linguaggio di programmazione Go

 Il linguaggio Go è detto "general purpose" vale a dire che è ideale per
 sviluppare applicazioni in un contesto generico e non specializzato, anche se
 ha trovato campo d'impiego soprattutto nel Web e nei sistemi enterprise.

 Go è un linguaggio fortemente tipizzato, cioè impone restrizioni sulle
 operazioni ammesse tra variabili di tipi diverso. Ha un garbage collector
 per la "pulizia" della memoria da tutte le strutture dati inutilizzate ed è
 predisposto per la programmazione concorrente. Un programma è costituito da
 pacchetti e moduli che permettono una gestione efficiente delle dipendenze.

 Innanzitutto il linguaggio Go è compilato, vale a dire che il codice sorgente
 è tradotto, con le sue dipendenze, da un compilatore in un file archivio, detto
 oggetto e poi un altro programma detto linker legge questo file e produce
 l'eseguibile.
 Queste operazioni sono combinate in un unico comando chiamato go che ha delle
 opzioni per compiere sottocomandi. Per esempio, il più semplice è run che
 effettua sia la compilazione di uno o più file di codice sorgente che il loro
 collegamento e poi, se il codice è sintatticamente corretto, esegue il
 programma eseguibile. In questo caso, il file oggetto e l'eseguibile sono
 creati temporaneamente e al termine dell'esecuzione vengono cancellati. Invece,
 qualora occorresse avere l'eseguibile permamentemente disponibile sul computer
 per rieseguirlo o redistribuirlo, allora il sottocomando build produce il
 programma nella stessa directory del file di codice sorgente.

 Il codice sorgente di un programma scritto in Go è contenuto in un file con
 estensione .go mentre l'eseguibile, se non diversamente specificato, ha lo
 stesso nome ma senza l'estensione .go.

 Il seguente codice Go è la versione del classico programma di saluto
 "Hello World" presentata all'inizio del libro "The C Programming Language"
 pubblicato nel 1978. Nella sua semplicità si possono comunque trovare alcune
 idee centrali del linguaggio di programmazione.
*/

/*
Tipicamente un programma scritto in linguaggio Go è una collezione di uno o più
pacchetti, conservati in directory separate, ciascuno dei quali consiste a sua
volta di uno o più file di codice sorgente. Ciò che contraddistingue un file
sorgente di un pacchetto da un altro è la "dichiarazione di pacchetto" che è la
prima istruzione, essa inizia con la parola chiave package seguita dal nome del
pacchetto. Si osserva che la specifica del linguaggio non impone di mantenere i
pacchetti in directory diverse. In particolare, la dichiarazione package main
ha un significato speciale per il compilatore: definisce un programa eseguibile
autosufficiente che è ben diverso da un pacchetto di libreria.
*/
package main

/*
È impensabile dover scrivere il codice Go che compie la stessa operazione in
ogni programma: se il codice è riutilizzabile allora conviene ridistribuirlo.
Go ha una libreria detta standard con oltre 100 pacchetti per compiere le
operazioni più comuni: input/output, ordinamento, ricerca, ecc. I pacchetti sono
stati scritti dagli scienziati di Google e dai programmatori del linguaggio Go.
Per poterli utilizzare è sufficiente importarli con l'istruzione import seguita
dal nome della libreria. In questo caso, il programma Hello utilizza il codice
per stampare sul terminale del computer che è contenuto nella libreria standard
fmt.

In generale, l'istruzione import dice al compilatore quali pacchetti sono
necessari per compilare il codice sorgente del programma. Questa istruzione deve
seguire la dichiarazione di pacchetto e comparire prima di ogni altra
dichiarazione di costante, variabile, tipo o funzione (l'ordine non ha
importanza), inoltre occorre importare solo i pacchetti strettamente necessari
al programma: la compilazione non viene portata a termine se ci sono pacchetti
inutilizzati o riferimenti a pacchetti non importati con questa istruzione.
*/
import "fmt"

/*
* La funzione main()

Nel pacchetto principale deve essere definita la funzione main() che è il punto
dove inizia l'esecuzione del programma.

In generale, una dichiarazione di funzione ha la seguente forma: inizia con
la parola chiave func, seguita dal nome di funzione con una lista di parametri
racchiusa da parentesi tonde, eventualmente una lista di risultati e infine il
corpo della funzione delimitato da parentesi graffe.
*/
func main() {
	// Non è necessario utilizzare punti e virgola per demarcare la fine di una
	// dichiarazione, in quanto i caratteri di fine linea vengono
	// automaticamente convertiti dal compilatore in un punto e virgola.
	// Questa regola vale quando ogni istruzione coincide con una linea di
	// codice; se invece due o più comandi compaiono sulla stessa riga allora
	// devono necessariamente essere separate da un punto e virgola.
	// Quindi la posizione dei caratteri di nuova linea è rilevante al fine di
	// ottenere il parsing corretto del codice Go.
	//
	// Per esempio, la parentesi graffa { deve essere inserita dopo la lista dei
	// risultati (o dopo la parentesi tonda se la funzione è main) e non può
	// trovarsi su una linea a se stante altrimenti il compilatore inserendo un
	// punto e virgola al posto del carattere di nuova linea, produce codice
	// sintatticamente scorretto.
	//
	// Go supporta il set di caratteri Unicode e pertanto è possibile utilizzare
	// ulteriori simboli che non compaiono nella codifica ASCII, per esempio
	// lettere di lingue straniere diverse dall'inglese, simboli pittografici,
	// emoticons, ecc.
	// La seguente istruzione crea una variabile il cui nome corrisponde alla
	// parola coreana mondo e allo stesso tempo memorizza un testo di saluto.
	var 세계 string = "Hello 세계! I'm Unicode"

	// La seguente riga di codice stampa la frase salvata nella variabile 세계
	// Se il terminale supporta Unicode, il programma stampa la frase
	// Hello 세계! I'm Unicode.
	//
	// La funzione Println() fa parte della libreria standard del linguaggio,
	// Stampa uno o più valori separati da spazi e con un carattere di nuova
	// linea dopo l'ultimo valore ricevuto.
	fmt.Println(세계)
}

/*
Il programma gofmt

Il linguaggio Go è molto rigido per quanto riguarda la formattazione del codice
sorgente, nel senso che si deve rispettare lo stile stabilito da Go. Per questo
motivo è stato creato il programma gofmt che formatta il codice conformemente
al formato standard prestabilito per il linguaggio. C'è anche l'opzione fmt per
il compilatore che invoca il comando gofmt su tutti i file di codice sorgente
presenti nella directory di lavoro (questo è il comportamento tipico).

Gli editor di testo che supportano il linguaggio Go vengono configurati per
eseguire gofmt ogni volta che si salva il file.

C'è un altro comando che facilita la scrittura dei programmi: goimorts aggiorna
le linee di codice che importano i pacchetti, aggiungendo quelli mancanti
quando ci sono dei riferimenti da qualche parte nel codice, oppure rimuovendo
gli import che sono inutilizzati. Goimports può anche formattare il codice
come gofmt per cui si ha con un unico comando due azioni estremamente utili.
*/
