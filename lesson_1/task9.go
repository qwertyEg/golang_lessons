package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    index := 2

    slice = append(slice[:index], slice[index+1:]...)
    fmt.Println("Слайс после удаления элемента:", slice)
}