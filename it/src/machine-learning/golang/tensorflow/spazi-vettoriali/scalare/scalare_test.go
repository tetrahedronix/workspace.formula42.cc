package main

import (
	"testing"

	tf "github.com/wamuir/graft/tensorflow"
)

func TestMoltPerScalare(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float32
		expected float32
	}{
		{"Scalari positivi", 3.0, 2.0, 6.0},
		{"Moltiplicazione con zero", 3.0, 0.0, 0.0},
		{"Moltiplicazione con negativo", 3.0, -4.0, -12.0},
		{"Moltiplicazione per l'unità", 4.4, 1.0, 4.4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tensorA, err := tf.NewTensor(tc.a)

			if err != nil {
				t.Fatalf("Impossibile creare tensorA: %v", err)
			}
			tensorB, err := tf.NewTensor(tc.b)

			if err != nil {
				t.Fatalf("Impossibile creare tensorB: %v", err)
			}

			resultTensor, err := MoltPerScalare(tensorA, tensorB)

			if err != nil {
				t.Fatalf("Errore in MoltPerScalare: %v", err)
			}

			if resultTensor == nil {
				t.Fatal("MoltPerScalare ha restituito un tensore nullo")
			}

			val, ok := resultTensor.Value().(float32)

			if !ok {
				t.Fatalf("Valore del tensore non è float32, ottenuto %T", val)
			}

			if val != tc.expected {
				t.Errorf("Prodotto atteso %v, ottenuto %v", tc.expected, val)
			}
		})
	}

}
