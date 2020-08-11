package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var DEBUG_FLAG bool = false

func main() {
	/*
	 * With combinations, order does not matter, and there cannot be
	 * repetition. For instance, if you pick 3 billiard balls out of a
	 * bag, and there are 16 total, the number of possible choices is 560.
	 *
	 *  Formula:
	 *
	 *        n!
	 *  --------------
	 *    r!(n - r)!
	 */

	blurb()

	n, r := parse_args(os.Args[1], os.Args[2])

	nb := bang(n)

	debug(fmt.Sprintf("mnb! (%d!) = %d", n, nb))

	rb := bang(r)

	debug(fmt.Sprintf("rb! (%d!) = %d", r, rb))

	nmr := big.NewInt(0)

	nmr.Sub(n, r)
	debug(fmt.Sprintf("n - r (%d - %d) = %d", n, r, nmr))

	nmrb := bang(nmr)

	debug(fmt.Sprintf("nmr! (%d!) = %d", nmr, nmrb))

	denom := big.NewInt(1)
	denom.Mul(rb, nmrb)

	debug(fmt.Sprintf("nb = %d", nb))
	debug(fmt.Sprintf("denom = %d", denom))

	//var ans, rem *big.Int
	//ans, rem = ans.DivMod(nb, denom, rem)
	//
	ans := nb.Div(nb, denom)

	fmt.Printf("For %d elements, taken %d at a time, ignoring order and without\nrepetition, you have %d possible combinations.\n", n, r, ans)

}

func blurb() {

	fmt.Println("----------------------------------------------------------------\n")
	fmt.Println("With combinations, order does not matter. For instance, if you")
	fmt.Println("pick 3 out of * 16 billiard balls, the number of possible")
	fmt.Println("choices is 560.\n")
	fmt.Println("  Formula:\n")
	fmt.Println("        n!")
	fmt.Println("  --------------")
	fmt.Println("    r!(n - r)!")
	fmt.Println("----------------------------------------------------------------\n")
}

// Brute-force factorial using recursion
// Uses math.big for big honkin' numbers
func bang(n *big.Int) *big.Int {

	one := big.NewInt(1)
	zero := big.NewInt(0)

	if n.Cmp(zero) == 0 {
		return one
	} else {
		t := big.NewInt(0)

		t.Sub(n, one)

		return t.Mul(n, bang(t))
	}
}

func parse_args(a, b string) (ni, ri *big.Int) {
	ERR := "\033[1;31m"
	NRM := "\033[0m"

	n, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Printf("%s%s\n%s%s\n", ERR, "Invalid value for n", err, NRM)
		os.Exit(1)
	}

	r, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		fmt.Printf("%s%s\n%s%s\n", ERR, "Invalid value for r", err, NRM)
		os.Exit(1)
	}

	// Make sure n > r
	if n <= r {
		fmt.Printf("%sError, n (%d) must be greater than r (%d)%s\n", ERR, n, r, NRM)
		os.Exit(1)
	}

	ni = big.NewInt(n)
	ri = big.NewInt(r)

	return
}

func debug(msg string) {
	BLCK := "\033[1;30m"
	NORM := "\033[0m"

	if DEBUG_FLAG {
		fmt.Printf("%s[ DEBUG: %s ]%s\n", BLCK, msg, NORM)
	}
}
