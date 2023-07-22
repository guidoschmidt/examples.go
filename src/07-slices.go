package main

import "fmt"

func main() {
	// Allocate memory for a 'slice' of three strings
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("length:", len(s))

	// Append additional strings
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appendd:", s)

	// Create a copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:2]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2:", l)

	l = s[2:]
	fmt.Println("slice 3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("temp:", t)

	matrix := make([][]int, 3)
	fmt.Println("2D size:", len(matrix))
	for i := 0; i < 3; i++ {
		inner_len := i + 1
		matrix[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("2D:", matrix)
}
