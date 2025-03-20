package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenCount := 0
    oddCount := 0

    for _, num := range slice {
        if num%2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }

    fmt.Printf("Четные: %d, Нечетные: %d\n", evenCount, oddCount)
}