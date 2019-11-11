package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Pick a number n > 1. Apply the following process repeatedly until n becomes 1.

If n is even, divide it by 2.
If n is odd, multiply it by 3, then add 1.
Return the number of steps needed.

In case you need the remainder of a division, use “a % b”.
*/

func collatz(number int) int {
	count := 0

	for number > 1 {
		if number%2 == 0 {
			number = number / 2
			count++
		} else {
			number = number*3 + 1
			count++
		}
	}
	return count
}
func main() {
	var n int
	var err error
	if len(os.Args) > 1 { // Read the number from the command line
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // Read the number interactively
		fmt.Println("Input a number > 1: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if n <= 1 {
		fmt.Println("n must be larger than 1.")
		return
	}
	c := collatz(n)
	// c := 0

	fmt.Println(n, "requires", c, "steps to reach 1.")
}
