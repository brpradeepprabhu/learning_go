package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// breaking the for loop
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	// continue for loop
	for i := 0; i < 10; i++ {
		if i < 5 {
			continue
		}
		fmt.Println("cont ", i)
	}
	//replace of while loop
	a, max := 5, 10
	for a < max {
		fmt.Println("max", a)
		a++
	}

	//infinite loop
	a = 0
	for {
		if a == 2 {
			break
		}
		fmt.Println("infinite", a)
		a++
	}
}
