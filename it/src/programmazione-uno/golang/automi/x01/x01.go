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
// Questo programma prende le parole in ingresso dagli argomenti della riga di
// comando e utilizza una macchina a stati finiti per determinare se ogni parola
// è accettata o respinta. Lo stato della macchina è rappresentato dalla
// variabile stato, che inizia a 0 e può transitare a 1 o 2 quando
// l'input è rispettivamente "0" o "1". Se l'input contiene altri caratteri, il
// programma stampa un messaggio di errore ed esce.

// Se lo stato finale della macchina è 2, il programma stampa un messaggio che
// indica che la parola è accettata. In caso contrario, stampa un messaggio che
// indica che la parola è respinta.
package main

import (
	"fmt"
	"os"
)

func main() {
	state := 0 // initial state

	for _, input := range os.Args[1:] {
		for _, char := range input {
			switch char {
			case '0':
				if state == 0 {
					state = 1 // go from state q0 to q1 with input 0
				} else if state == 2 {
					state = 1 // go from state q2 to q1 with input 0
				}
			case '1':
				if state == 1 {
					state = 2
				} else if state == 2 {
					state = 0
				}
			default:
				fmt.Println("invalid input:", string(char))
				return
			}
		}

		if state == 2 {
			fmt.Println("accepted:", input)
		} else {
			fmt.Println("rejected:", input)
		}
	}
}
