package main

import "fmt"

func main() {
    arr := [8]int{5, 3, 8, 1, 9, 4, 7, 2}
    num := 9
    index := -1

    for i, value := range arr {
        if value == num {
            index = i
            break
        }
    }

    fmt.Println("Индекс элемента:", index)
}