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
