package main

import "fmt"

func main() {
    arr := [7]int{3, 7, 2, 9, 5, 1, 6}
    max := arr[0]

    for _, num := range arr {
        if num > max {
            max = num
        }
    }

    fmt.Println("Максимальный элемент:", max)
}