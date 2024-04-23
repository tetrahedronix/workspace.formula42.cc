package main

import (
	"bufio"
	"bytes"
	"testing"
)

// TestSearchPattern esegue i test sulla funzione searchPattern per verificare
// se accetta correttamente le parole che contengono tutte e cinque le vocali
// in ordine lessicografico.
func TestSearchPattern(t *testing.T) {
	tests := []struct {
		word     string // Parola da testare
		expected bool   // Valore atteso restituito dalla funzione searchPattern
	}{
		{"", false},
		{"abstemious", true}, // Testa una parola che contiene tutte e cinque le vocali
		{"aeiou", true},      // Testa una parola che contiene tutte e cinque le vocali in ordine
		{"aeioua", true},     // Testa una parola che contiene tutte e cinque le vocali in ordine
		{"aaeioou", true},
		{"aaeeiioouu", true},
		{"facetious", true},     // Testa un'altra parola che contiene tutte e cinque le vocali
		{"ae", false},           // Testa una parola che non contiene tutte e cinque le vocali in ordine
		{"aieiou", false},       // Testa una parola che contiene tutte e cinque le vocali in ordine
		{"iou", false},          // Testa una parola che non contiene tutte e cinque le vocali in ordine
		{"sacrilegious", false}, // Testa un'altra parola che contiene tutte e cinque le vocali
		{"uoieia", false},       // Testa una parola che contiene tutte le vocali in ordine opposto al lessicografico
		{"zoo", false},          // Testa una parola che non contiene tutte e cinque le vocali in ordine
	}

	// Cicla su ogni test
	for _, test := range tests {
		// Esegue un sotto-test per ciascun test
		t.Run(test.word, func(t *testing.T) {
			// Converte la parola in un buffer di byte
			buffer := []byte(test.word)
			// Crea un reader basato sul buffer
			reader := bytes.NewReader(buffer)
			// Crea uno scanner di parole
			wordScanner := bufio.NewScanner(reader)
			wordScanner.Split(bufio.ScanWords)
			// Scansiona la parola
			wordScanner.Scan()
			word := wordScanner.Text()

			// Chiama la funzione searchPattern e ottiene il risultato
			//result := searchPattern(word)
			result := searchPattern(word)
			// Verifica se il risultato ottenuto corrisponde all'aspettativa
			if result != test.expected {
				t.Errorf("Per la parola '%s', ci si aspettava %t ma è stato ottenuto %t", test.word, test.expected, result)
			}
		})
	}
}
