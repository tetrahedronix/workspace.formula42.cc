package main

import (
	"testing"
)

// Struttura di test che contiene le parole da testare
var testData = []struct {
	word        string // Parola da testare
	shouldBeOK  bool   // Indica se ci si aspetta che la parola sia accettata
	shouldBeOKB bool   // Indica se ci si aspetta che la parola sia accettata da searchPatternB
}{

	{"01100100", false, true},
	{"10000000", false, true},
	{"11111110", false, false},
	{"00001000", false, true},
	{"101010", false, true},     // Parola con numero pari di 1
	{"100000000", false, true},  // Parola con 8 bit
	{"0101010111", true, false}, // Parola con 8 bit ma numero dispari di 1
	{"101011", true, true},      // Parola con numero dispari di 1
	{"00000000", true, true},
	{"00000011", true, true},
	{"01010101", true, true},
	{"00000", true, true}, // Parola con solo 0
	{"", true, true},      // Parola vuota

}

// TestSearchPatternA verifica il corretto funzionamento della funzione searchPatternA
func TestSearchPatternA(t *testing.T) {
	for _, test := range testData {
		result := searchPatternA(test.word)
		if result != test.shouldBeOK {
			t.Errorf("Expected searchPatternA('%s') to return %t, but got %t", test.word, test.shouldBeOK, result)
		}
	}
}

// TestSearchPatternB verifica il corretto funzionamento della funzione searchPatternB
func TestSearchPatternB(t *testing.T) {
	for _, test := range testData {
		result := searchPatternB(test.word)
		if result != test.shouldBeOKB {
			t.Errorf("Expected searchPatternB('%s') to return %t, but got %t", test.word, test.shouldBeOKB, result)
		}
	}
}
