package main

import (
	"fmt"    // Importa il pacchetto fmt per l'output formattato
	"regexp" // Importa il pacchetto regexp per lavorare con le espressioni regolari
)

func main() {
	// Definisce un'espressione regolare per gli indirizzi email utilizzando la funzione MustCompile
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Crea una lista di indirizzi email di esempio da verificare
	testEmails := []string{
		"user@example.com",            // Email valida
		"user.name+tag@example.co.uk", // Email valida con simboli e dominio a più livelli
		"invalid.email@com",           // Email invalida (estensione di dominio troppo corta)
		"12345@domain.io",             // Email valida composta da numeri
		"adamdomain.io",               // Email invalida: manca @
		"@adam@adam.io",               // Sebbene invalida, viene accettata
	}

	// Itera attraverso la lista di indirizzi email
	for _, email := range testEmails {
		// Verifica se l'email corrisponde all'espressione regolare
		if emailRegex.MatchString(email) {
			// Stampa se l'indirizzo email è valido
			fmt.Printf("%s è un indirizzo email valido\n", email)
		} else {
			// Stampa se l'indirizzo email non è valido
			fmt.Printf("%s non è un indirizzo email valido\n", email)
		}
	}
}
