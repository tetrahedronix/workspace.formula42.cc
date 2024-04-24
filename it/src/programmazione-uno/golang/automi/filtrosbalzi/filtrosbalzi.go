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
    HEREUNDER.package main
*/
package main

import (
	"fmt"
)

func checkInput(s string, r rune) {

}

func automaBounceFilter(w string) bool {

	currentState := 'a' // stato iniziale
	for _, char := range w {
		// Esegue la transizione in base allo stato corrente e al carattere letto
		switch currentState {
		case 'a':
			if char == '0' {
				currentState = 'a'
				fmt.Print("0")
			} else if char == '1' {
				currentState = 'b'
				fmt.Print("0")
			}
		case 'b':
			if char == '0' {
				currentState = 'a'
				fmt.Print("0")
			} else if char == '1' {
				currentState = 'c'
				fmt.Print("1")
			}
		case 'c':
			if char == '0' {
				currentState = 'd'
				fmt.Print("1")
			} else if char == '1' {
				currentState = 'c'
				fmt.Print("1")
			}
		case 'd':
			if char == '0' {
				currentState = 'a'
				fmt.Print("0")
			} else if char == '1' {
				currentState = 'c'
				fmt.Print("1")
			}
		}
	}
	// Verifica se lo stato finale è accettante
	fmt.Println()
	return currentState == 'c'

}

func main() {

	var words = []string{
		"",
		"0",
		"01",
		"010",
		"0101",
		"1010",
		"0011",
		"01011",
		"010110",
		"0101101",
	}

	for _, str := range words {
		if automaBounceFilter(str) {
			fmt.Printf("'%s' OK\n", str)
		} else {
			fmt.Printf("'%s' NO\n", str)
		}
	}
}
