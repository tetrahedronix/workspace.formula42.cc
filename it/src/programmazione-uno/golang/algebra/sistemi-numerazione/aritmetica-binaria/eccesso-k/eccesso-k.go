// Implementazione di un codice in eccesso-k
//
// # Descrizione del problema
//
// Si richiede di implementare un programma che converta un numero intero N in
// una rappresentazione binaria utilizzando il codice in eccesso-k. Il programma
// dovrà garantire che la rappresentazione finale utilizzi esattamente n bit.
//
// Input:
//
// Il programma riceverà in input tre valori:
// - Un intero positivo k che rappresenta il valore di eccesso da utilizzare
// - Un intero positivo n che rappresenta il numero di bit da utilizzare nella rappresentazione finale
// - Un intero N che rappresenta il numero da convertire
//
// Output:
//
// Il programma dovrà stampare una stringa composta da n caratteri binari
//
//	(0 e 1) che rappresenta il numero N codificato in eccesso-k.
//
// Vincoli:
//
// - 0 ≤ k ≤ 1000
// - 1 ≤ n ≤ 32
// - -10^9 ≤ N ≤ 10^9
// - Assicurarsi che N + k sia rappresentabile con n bit (ovvero: 0 ≤ N + k < 2^n)
//
// # Esempio 1
//
// Input: 3 4 5
// Output:1000
//
// Spiegazione: Con k=3, n=4 e N=5, il valore da rappresentare è 5+3=8, che in
// binario su 4 bit è 1000.
//
// # Esempio 2
//
// Input: 7 5 -2
// Output: 00101
//
// Spiegazione: Con k=7, n=5 e N=-2, il valore da rappresentare è -2+7=5, che in
// binario su 5 bit è 00101.
//
// # Esempio 3
//
// Input: 10 6 22
// Output: 100000
//
// Spiegazione: Con k=10, n=6 e N=22, il valore da rappresentare è 22+10=32,
// che in binario su 6 bit è 100000.
//
// Note:
//
//   - Se N+k non è rappresentabile con n bit, il programma dovrebbe generare un
//     messaggio di errore appropriato.
//   - La rappresentazione binaria deve sempre utilizzare esattamente n bit,
//     aggiungendo zeri non significativi a sinistra se necessario.
//   - Per numeri negativi, occorre prima sommare k e poi convertire il risultato
//     in binario (come nell'esempio 2).
//
// Suggerimenti di implementazione:
//
// 1. Calcolare il valore M = N + k
// 2. Verificare che 0 ≤ M < 2^n
// 3. Convertire M in binario
// 4. Assicurarsi che la rappresentazione binaria abbia esattamente n bit
package main

import (
	"fmt"
	"math"
)

type EccessoK struct {
	k, n, N int
	bin     string
}

func newEccessoK(k, n, N int) *EccessoK {
	e := &EccessoK{k, n, N, ""}

	// Verifica che N+k sia rappresentabile con n bit
	if e.N+e.k >= 0 && e.N+e.k < int(math.Pow(2, float64(e.n))) {
		// Converti N+k direttamente in binario con zeri a sinistra usando Sprintf
		e.bin = fmt.Sprintf("%0*b", e.n, N+k)
	}

	return e
}

func (e *EccessoK) GetBin() string {
	return e.bin
}

func (e *EccessoK) String() string {
	if e.bin == "" {
		return "Errore: N+k non è rappresentabile con n bit"
	}

	return fmt.Sprintf("eccesso-%d: %s", e.k, e.bin)
}

func main() {
	var k, n, N int

	fmt.Scan(&k, &n, &N)

	ek := newEccessoK(k, n, N)

	fmt.Println(ek)
}
