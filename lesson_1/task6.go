package main

import "fmt"

func main() {
    slice := []int{3, 7, 2, 9, 5, 1, 6}
    max := slice[0]

    for _, num := range slice {
        if num > max {
            max = num
        }
    }

    fmt.Println("Максимальный элемент:", max)
}