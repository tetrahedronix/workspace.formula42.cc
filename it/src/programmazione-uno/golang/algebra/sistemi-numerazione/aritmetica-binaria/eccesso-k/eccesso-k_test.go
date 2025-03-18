package main

import (
	"testing"
)

// Una funzione di test che verifica la corretta conversione di un numero in eccesso-k
func TestEccessoKBasic(t *testing.T) {
	// caso di test: k=3, n=4, N=5
	ek := newEccessoK(3, 4, 5)
	expected := "1000"

	if ek.GetBin() != expected {
		t.Errorf("Errore: atteso %s, ottenuto %s", expected, ek.GetBin())
	}
}

// Una funzione che verifica diversi casi di test utilizzando una tabella di test
func TestEccessoKMultipleCases(t *testing.T) {
	testCases := []struct {
		k, n, N  int
		expected string
	}{
		{3, 4, 5, "1000"},
		{7, 5, -2, "00101"},
		{10, 6, 22, "100000"},
		{0, 3, 4, "100"},
		{1, 4, 0, "0001"},
	}

	for _, tc := range testCases {
		ek := newEccessoK(tc.k, tc.n, tc.N)
		if ek.GetBin() != tc.expected {
			t.Errorf("Errore: atteso %s, ottenuto %s", tc.expected, ek.GetBin())
		}
	}
}

// Un test per verificare il comportamento quando N+k non è rappresentabile in n bit
func TestEccessoKOutOfRnage(t *testing.T) {
	// Caso di test dove N+k non è rappresentabile con n bit
	ek := newEccessoK(3, 2, 5) // 5+3=8, che richiede 4 bit

	if ek.GetBin() != "" {
		t.Errorf("Errore: atteso \"\", ottenuto %s", ek.GetBin())
	}
}

// Test per verificare che la lunghezza della stringa binaria sia sempre esattamente n bit
func TestEccessoKBinaryLength(t *testing.T) {
	testCases := []struct {
		k, n, N int
	}{
		{3, 8, 5},
		{7, 10, -2},
		{10, 12, 22},
	}

	for _, tc := range testCases {
		ek := newEccessoK(tc.k, tc.n, tc.N)
		if len(ek.GetBin()) != tc.n {
			t.Errorf("Errore: atteso %d bit, ottenuto %d bit", tc.n, len(ek.GetBin()))
		}
	}

}
