package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    sum := 0

    for _, num := range arr {
        sum += num
    }

    fmt.Println("Сумма элементов массива:", sum)
}