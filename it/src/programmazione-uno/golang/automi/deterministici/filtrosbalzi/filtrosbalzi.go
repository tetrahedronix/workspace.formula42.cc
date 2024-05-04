/*
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

/*
#include <stdio.h>

int getChar() {
	return getchar();
}
*/
import "C"

import (
	"fmt"
	"os"
)

func gotoQ0() {

	fmt.Print("0")
	r := rune(C.getChar())

	// Transita allo stato q0
	if r == '0' {
		gotoQ0()
	}

	// Transita allo stato q1
	if r == '1' {
		gotoQ1()
	}

	// Transita allo stato finale q4
	gotoQ4()
}

func gotoQ1() {

	fmt.Print("0")
	r := rune(C.getChar())

	// Transita allo stato q0
	if r == '0' {
		gotoQ0()
	}

	// Transita allo stato q2
	if r == '1' {
		gotoQ2()
	}

	// Transita allo stato finale q4
	gotoQ4()
}

func gotoQ2() {

	fmt.Print("1")
	r := rune(C.getChar())

	// Transita allo stato q3
	if r == '0' {
		gotoQ3()
	}

	// Transita allo stato q2
	if r == '1' {
		gotoQ2()
	}

	// Transita allo stato finale q4
	gotoQ4()
}

func gotoQ3() {

	fmt.Print("1")
	r := rune(C.getChar())

	// Transita allo stato q0
	if r == '0' {
		gotoQ0()
	}

	// Transita allo stato q2
	if r == '1' {
		gotoQ2()
	}

	// Transita allo stato finale q4
	gotoQ4()

}

func gotoQ4() {

	os.Exit(1)
}

func main() {

	gotoQ0()
	fmt.Println()
}
