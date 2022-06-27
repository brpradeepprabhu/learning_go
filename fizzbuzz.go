package main

import "fmt"

/*
 Fizz buzz
if number is divisible by 3 then we want to display fizz
if number is divisible by 5 then we want tod display buzz
if number is divisible by both then we  want to display fizz buzz
*/
func main() {
	for i := 0; i < 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
