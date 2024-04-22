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
