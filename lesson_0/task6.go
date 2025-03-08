package main

import "fmt"

func main() {
    s := "Привет, мир!"
    count := 0
    for range s {
        count++
    }
    fmt.Println(count)
}

//Программа использует цикл for range, чтобы перебрать строку по рунам.
//Переменная count увеличивается на каждой итерации, подсчитывая количество символов.
//Результат: 12 (символов в строке "Привет, мир!"). Пробел это тоже символ!!!!