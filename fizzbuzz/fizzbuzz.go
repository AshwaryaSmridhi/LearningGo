package main

import (
	"fmt"
	"os"
	"strconv"
)

/*Write a program that plays the game of FizzBuzz. The rules are simple:

Count from 1 to a given number n.
Print out each number, with the following exceptions:
If the number is divisible by 3, print “Fizz” instead of the number.
If the number is divisible by 5, print “Buzz”.
If the number is divisible by 15, print “FizzBuzz”.
The code below reads a number from the command line (default is 50 if no number is given) and calls function fizzbuzz with that number.

Fill the body of fizzbuzz, using a for loop and a switch statement.*/

func fizzbuzz(n int) {
	i := 1
	for i <= 50 {
		// if i%15 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if i%5 == 0 {
		// 	fmt.Println("Buzz")
		// } else if i%3 == 0 {
		// 	fmt.Println("Fizz")
		// } else {
		// 	fmt.Println(i)
		// }

		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
		i = i + 1
	}
}
func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
