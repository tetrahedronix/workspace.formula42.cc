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
// x01y è un automa che accetta tutte e sole le stringhe di 0 e di 1 che
// contengono la sequenza 01 da qualche parte nella parola fornita al programma
// sulla riga di comando.
package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "word")
		os.Exit(1)
	}

	var state int64 = 0
	word := os.Args[1]

	for _, c := range word {
		state = x01y(state, c)
	}

	if state == 2 {
		fmt.Println("Accepted: ", word)
	} else {
		fmt.Println("Rejected: ", word)
	}
}

func x01y(s int64, c rune) int64 {

	switch s {
	case 0:
		if c == '0' {
			return 1
		}
		if c == '1' {
			return 0
		}
	case 1:
		if c == '0' {
			return 1
		}
		if c == '1' {
			return 2
		}
	case 2:
		if c == '0' || c == 1 {
			return 2
		}
	}

	return s
}
