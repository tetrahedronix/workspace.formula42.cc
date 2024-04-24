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
init stampa la stringa golang.org e il valore unicode del carattere b.
*/
package main

import "fmt"

var glob string

/*
init assegna la stringa "golang.org" alla variabile globale glob
*/
func init() {
	// Le variabili possono essere inizializzate usando la funzione init()
	// dichiarata nel blocco del pacchetto.
	glob = "golang.org"
}

/*
init questa funzione init non fa nulla.
È del tutto legale definire multiple funzioni init() per pacchetto, anche
dentro il singolo file sorgente; inoltre, come qualsiasi altra funzione
anche init() può essere vuota, ossia priva di istruzioni.
*/
func init() {}

func init() {}

/*
main stampa i valori di glob e c
*/
func main() {
	// Nel blocco pacchetto, l'identificatore init può essere utilizzato solo per
	// dichiarare le funzioni init, ma l'identificatore stesso non è dichiarato.
	// Pertanto non è possibile fare riferimento alle funzioni init da qualsiasi
	// altra parte del programma.
	// init() // undefined init

	// glob è stata inizializzata precedentemente con la stringa golang.org da
	// una funzione init() che ha precedenza su main()
	fmt.Println(glob, c)
}

/*
L'intero pacchetto viene inizializzato assegnando valori iniziali a tutte le
sue variabili a livello di pacchetto seguite dalla chiamata di tutte le funzioni
init nell'ordine in cui appaiono nel sorgente, possibilmente in più file, come
presentate al compilatore. Nel seguente spezzone di codice, c è una variabile
dichiarata a livello di pacchetto (comparendo in questo file sorgente con la
lettera minuscola), con un valore iniziale 0 che è il default per questo tipo di
dato e la funzione init() la inizializza con il valore 98.
*/
var c rune

func init() {

	c = 'b'
}
