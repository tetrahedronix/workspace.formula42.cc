package main

import "testing"

// TestSearchPronom è una funzione di test per la funzione searchPronom.
func TestSearchPronom(t *testing.T) {
	// Definizione dei casi di test con input e output attesi.
	testCases := []struct {
		input    string // input è la parola da testare.
		expected bool   // expected indica se ci si aspetta che la parola sia un pronome o meno.
	}{
		{"he", true},
		{"his", true},
		{"him", true},
		{"she", true},
		{"her", true},
		{"hers", true},
		{"hersa", false},
		{"ers", false},
		{"h", false},
		{"s", false},
		{"hehe", false},
		{"shehe", false},
		{"himmm", false},
		{"hhee", false},
	}

	// Itera attraverso i casi di test.
	for _, tc := range testCases {
		// Esegue un sottotest per ciascun caso di test.
		t.Run(tc.input, func(t *testing.T) {
			// Chiama la funzione searchPronom con l'input del caso di test e controlla se il risultato è quello atteso.
			if got := searchPronom(tc.input); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
