# Problem Definition

➜ Given an array of integers <br>
➜ Given an integer target value <br>
➜ Return the indices of the 2 numbers that add up to that given target<br>

# Example

➜ [1, 3, 7, 9, 2], target = 11 <br>
➜ result should be = (3, 4) <br>

# Constraints

➜ All numbers are positive <br>
➜ There are no duplicates numbers in the array <br>
➜ There might not always be a solution for the problem <br>
➜ We can recieve an empty array <br>
➜ If there is no solution return null (depends on the language you use) <br>
➜ Only one pair of numbers will be the solution, in other means we can't have more than one solution

# Test cases

```bash
[5], target = 5
➜ result = null
```

```bash
[1, 2, 4, 9], target = 13
➜ result =  [2, 3]
```

```bash
[], target = 8
➜ result = null
```

# How i optimized the algorithm to take O(n) instead of O(n*n) time complexity 
```go
func trivial_solution(inp []int, target int) []int {
    // validation for requierments
    if len(inp) == 0 || len(inp) == 1 {
        return nil
	}

    // outer loop for picking a number
	for i := 0; i < len(inp); i++ {
        // inner loop for finding the number to add it to the pciked one to get the target
		for j := i + 1; j < len(inp); j++ {
            if inp[j] == target-inp[i] {
                return []int{i, j}
			}
		}
	}
	return nil
}
```
- the outer loop is used to pick up a number and based on it we try to find if there is any number in the list contains the other number which if its added to our picked number it will give the target

- the purpose of the inner loop is to find that number to add.
The only reason i need this loop is because the picked_number is live in the outer loop so we have to loop through all the rest of the array while we have the picked_number stored .. so can we store it in another way without having to pick it up and loop through the rest of numbers ! 

➜ **`Solution`**: 
```go
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
```
- pick a number using the outer loop
- define a hash-table
- check if the target-picked_number is not stored in the hash before
- if its not stored before, store it with the index of the picked_number 
- if its stored before, then return the index of the current picked_number, and the value where we find it stored in the hash-map <br>

    ```bash
    [1, 2, 4, 9], target = 13
    ➜ picked = 1, map = {}, num_to_find = 13-1 = 12, since map[picked] = map[1] = 0
       not stored before => store map[num_to_find] = index_of_picked_num => map={[12, 0]}
       this menas the value 12 is the num_to_find for number at index = 0.
       
    ➜ picked = 2, map = {[1, 0]}, num_to_find = 13-2 = 11 , since map[2] = 0
       not stored before => store map[num_to_find] = index_of_picked_num => map={[12, 0], [11, 1]}

    ➜ picked = 4, map = {[1, 0], [2, 1]}, num_to_find = 13-4 = 9, since map[4]=0 
       not stored before => map={[12, 0], [11, 1], [9, 2]}

    ➜ picked = 9, map = {}, num_to_find = 13-9 = 4, since map[9] is stored before 
       return the index_of_picked_num (index of 9) and the value fromt hash map of the map[9]
       return 3, 2
    ```
