package main

import "fmt"

func main() {
    arr := [6]int{1, 2, 3, 4, 5, 6}

    for i := 0; i < len(arr)/2; i++ {
        arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
    }

    fmt.Println("Реверс массива:", arr)
}