package main

func main() {
	arr := []int{16, 8, 42, 4, 23, 15}
	x := arr[0]
	for _, value := range arr[1:] {
		if x < value {
			x = value
		}
	}
	print("the maximal value", x)
}
