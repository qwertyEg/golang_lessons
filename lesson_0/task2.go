package main

import (
	"fmt"
	"strconv"
)

func main() {
    var input string
    fmt.Scanln(&input)
    num, _ := strconv.Atoi(input)
    fmt.Println(num)
}

// Программа считывает строку с помощью Scanln, 
// преобразует её в число с помощью Atoi и выводит результат.
//  Обратите внимание, что ошибка игнорируется (используется _), 
// но в реальном коде её нужно обрабатывать.