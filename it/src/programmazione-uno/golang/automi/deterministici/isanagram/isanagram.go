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

package main

import (
	"io"
	"os"
)

func main() {

}

// readFileToBuffer legge il contenuto di un file dato il suo nome e
// e restituisce il contenuto come un buffer di byte.
func readFileToBuffer(f string) ([]byte, error) {
	// Apre il file in lettura
	file, err := os.Open(f)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Legge il contenuto del file in un buffer
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}

// stringsFromBuffer legge i caratteri da un buffer di byte e chiama la
// funzione searchWords per trovare parole che sono composte da lettere che
// compaiono in un'altra parola specificata come secondo argomento.
func stringsFromBuffer(buffer []byte, input string) {

	// Converte il buffer di byte in una stringa
	txt := string(buffer)

	pos := s
}
