package main

import "fmt"

func findMissingNumbers(allNumbers, input []int) map[int]bool {
    present := make(map[int]bool)
    for _, num := range input {
        present[num] = true
    }
    
    missing := make(map[int]bool)
    for _, num := range allNumbers {
        if !present[num] {
            missing[num] = true
        }
    }
    
    return missing
}

func main() {
    allNumbers := []int{1, 2, 3, 4, 5}
    input := []int{2, 3, 5}
    fmt.Println(findMissingNumbers(allNumbers, input))
}