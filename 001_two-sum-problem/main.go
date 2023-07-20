package main

import "fmt"

func main() {
	inp := []int{1, 3, 9, 3, 5}
	target := 14
	fmt.Println("O(n*n) solution")
	fmt.Println(trivial_solution(inp, target))
	fmt.Println("O(n) solution")
	fmt.Println(optimized_solution(inp, target))
}

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

//	inp := []int{1, 3, 9, 3, 5}
func optimized_solution(inp []int, target int) []int {
	if len(inp) == 0 || len(inp) == 1 {
		return nil
	}
	store := make(map[int]int) // {}
	var num_to_find int        // = 0
	for idx, val := range inp {
		num_to_find = target - val // = 13
		// check if this value is the num_to_find for any other number in the given array
		if store[val] == 0 {
			// its not visited before
			// so store it in the map
			store[num_to_find] = idx
		} else {
			return []int{store[val], idx}
		}
	}
	return nil
}
