package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    sum := 0

    for _, num := range slice {
        sum += num
    }

    fmt.Println("Сумма элементов слайса:", sum)
}