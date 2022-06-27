package main

import "fmt"

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, len int) (*Square, error) {
	if len < 0 {
		return nil, fmt.Errorf("length is less than zero")
	}
	sq := Square{
		X: x, Y: y,
		Length: len,
	}
	return &sq, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	sq, err := NewSquare(1, 1, 10)
	if err != nil {
		fmt.Printf("Error in creating square %s", err)
	}
	fmt.Printf("%#v\n", sq)
	fmt.Printf("%+v\n", sq)
	sq.Move(2, 3)
	fmt.Println(sq.Area())
}
