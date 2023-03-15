package main

import "fmt"

func main() {
	// n1 := 2
	// s1 := []int{1, 0}
	// a1 := []int{2, 1}

	// n2 := 3
	// s2 := []int{0, 1, 1}
	// a2 := []int{1, 1, 1}

	n1 := 4
	s1 := []int{1, 0, 1, 0}
	a1 := []int{2, 3, 2, 3}

	n2 := 2
	s2 := []int{1, 0}
	a2 := []int{4, 6}

	var num1, num2, bits int

	for i := 0; i < n1; i++ {
		for j := 0; j < a1[i]; j++ {
			num1 = num1 << 1
			num1 += s1[i]
			bits++
		}
	}

	for i := 0; i < n2; i++ {
		for j := 0; j < a2[i]; j++ {
			num2 = num2 << 1
			num2 += s2[i]
		}
	}

	fmt.Printf("%08b\n", num1)
	fmt.Printf("%08b\n", num2)

	var years int
	fmt.Print("diff pos: ")
	for i := 0; i < bits; i++ {
		if ((num1 >> i) & 1) != ((num2 >> i) & 1) {
			fmt.Printf("%d, ", bits-i)
			years += bits - i
		}
	}

	fmt.Printf("\nnumber of years: %d\n", years)
}
