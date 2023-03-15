package main

import "fmt"

func task(n, m int) {
	var start, j int
	var b bool

	for i := 0; i < n; i++ {
		if i%m == 0 || i == 0 {
			if m%2 != 0 {
				start = m/2 + 1
				b = true
			} else {
				start = m / 2
				b = false
			}
			j = 1
			fmt.Println(start)
			continue
		}

		if b {
			fmt.Println(start - j)
			if m%2 == 0 {
				j++
			}
		} else {
			fmt.Println(start + j)
			if m%2 != 0 {
				j++
			}
		}

		b = !b
	}
}

func main() {
	// var n, m int
	// fmt.Scanf("%d %d", &n, &m)
	// fmt.Println(n, m)

	task(10, 5)
	fmt.Println()
	task(12, 6)
}
