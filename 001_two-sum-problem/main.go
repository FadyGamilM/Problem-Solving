package main

import "fmt"

func main() {
	inp := []int{3, 2, 4}
	target := 6
	fmt.Println("O(n*n) solution")
	fmt.Println(trivial_solution(inp, target))
	fmt.Println("O(n) solution")
	fmt.Println(optimized_solution(inp, target))
}

// inp := []int{1, 3, 9, 3, 5}
func optimized_solution(inp []int, target int) []int {
	if len(inp) == 0 || len(inp) == 1 {
		return nil
	}

	store := make(map[int]int)
	num_to_find := 0

	for idx, val := range inp {
		_, exists := store[inp[idx]]
		num_to_find = target - val
		if exists {
			// already exists, so return the result
			return []int{store[val], idx}
		} else {
			store[num_to_find] = idx
		}
	}
	return nil
}

// [3,2,4], target = 6

// num_to_find = 0
// store = {}

// idx = 0, val = 3
// ntf = 3
// store = {[3, 0]}

// val = 2
// ntf = 4
// store = {[3, 0], [4, 1]}

// val = 4
// ntf = 4
// store = {[3, 0], [4, 1]}

// return []int{0, 1}

// space complexity = O(1)
func trivial_solution(inp []int, target int) []int {
	if len(inp) == 0 || len(inp) == 1 {
		return nil
	}
	// time complexity = O(n*n)
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			if inp[j] == target-inp[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}
