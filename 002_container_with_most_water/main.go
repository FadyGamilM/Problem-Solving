package main

import "fmt"

func main() {
	fmt.Println(trivial_soln([]int{6, 9, 2, 5, 4}))
	fmt.Println(trivial_soln([]int{7, 1, 2, 3, 9}))

}

func trivial_soln(hs []int) []int {
	largest_area := 0

	selected_heights := []int{0, 0}

	for i := 0; i < len(hs); i++ {
		// initialize vars
		picked_height := hs[i]

		// looking for the combinator height
		for j := i + 1; j < len(hs); j++ {
			if picked_height < hs[j] {
				calc_area := picked_height * (j - i)
				if calc_area > largest_area {
					largest_area = calc_area
					selected_heights[0] = i
					selected_heights[1] = j
				}
			} else {
				calc_area := hs[j] * (j - i)
				if calc_area > largest_area {
					largest_area = calc_area
					selected_heights[0] = i
					selected_heights[1] = j
				}
			}
		}

	}

	return []int{selected_heights[0], selected_heights[1]}
}
