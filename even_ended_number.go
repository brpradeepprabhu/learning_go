package main

import "fmt"

/**
To find the first number and last number is same
increment the count
*/
func main() {
	count := 0
	for i := 1000; i < 9999; i++ {
		for j := i; j < 9999; j++ {
			s := fmt.Sprintf("%d", i*j)
			if s[0] == s[len(s)-1] {
				count += 1
			}
		}
	}
	fmt.Println("count", count)
}
