package main

import "testing"

func TestSearchPattern(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"nomeVariabile", true},
		{"lunghezzaStringa", true},
		{"valore_intero", true},
		{"array_di_caratteri", true},
		{"VAR", true},
		{"CONSTANTE", true},
		{"str", true},
		{"nomeVariabile_1", true},
		{"lunghezzaStringa_123", true},
		{"valore_intero_42", true},
		{"array_di_caratteri_999", true},
		{"Puntatore_a_struttura_007", true},
		{"VAR_3", true},
		{"CONSTANTE_42", true},
		{"str_abc", true},
		{"total3score", true},
		{"value4you", true},
		{"test123value", true},
		{"123var", false}, // Non può iniziare con un numero
		{"_underscore", false},
		{"3d-model", false}, // Non può iniziare con un numero
		{"-ok", false},      // Il trattino non è un carattere valido per un identificatore
		{"doub-", false},    // Il trattino non è un carattere valido per un identificatore
		{"$asin", false},    // Il dollaro non è un carattere valido per un identificatore
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if got := searchPattern(tc.input); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
