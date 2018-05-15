package main

import "fmt"

func eratosthenes(n int) []int {
	a := []int{}
	temp := make([]bool, n+1)
	for t := range temp {
		temp[t] = true
	}

	for i := 2; i < n+1; i++ {
		if temp[i] {
			j := i * i
			for j <= n {
				temp[j] = false
				j = j + i
			}
		}
	}
	for i := 2; i < n+1; i++ {
		if temp[i] {
			a = append(a, i)
		}
	}
	return a
}

func main() {
	fmt.Println(eratosthenes(120))
}
