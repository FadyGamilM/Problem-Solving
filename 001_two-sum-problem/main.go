package main

import "fmt"

func main() {
	inp := []int{1, 3, 9, 3, 5}
	target := 14
	fmt.Println(trivial_solution(inp, target))

}

func trivial_solution(inp []int, target int) []int {
	if len(inp) == 0 || len(inp) == 1 {
		return nil
	}
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			if inp[j] == target-inp[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}
