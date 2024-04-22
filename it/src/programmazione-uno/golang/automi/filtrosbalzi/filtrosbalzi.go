/**
 * Copyright (C) 2023 Giulio Carlo
 *
 * This file is part of programmazione-uno.
 *
 * programmazione-uno is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * programmazione-uno is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with programmazione-uno.  If not, see <http://www.gnu.org/licenses/>.
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
