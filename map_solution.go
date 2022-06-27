package main

import (
	"fmt"
	"strings"
)

func main() {
	text := ` Needles and pins 
Needles and pins
Saw me a sail 
catch the wind`
	count := map[string]int{}
	s := strings.Fields(text)

	for _, value := range s {
		count[strings.ToLower(value)]++
	}
	fmt.Println(count)
}
