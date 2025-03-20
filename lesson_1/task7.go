package main

import "fmt"

func main() {
    slice := []int{10, 4, 8, 2, 7, 3}
    min := slice[0]

    for _, num := range slice {
        if num < min {
            min = num
        }
    }

    fmt.Println("Минимальный элемент:", min)
}