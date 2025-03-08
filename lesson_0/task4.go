package main

import (
	"fmt"
	"strconv"
)

func main() {
    str := "1010"
    num, _ := strconv.ParseInt(str, 2, 64)
    fmt.Println(num)
}

//Программа преобразует двоичную строку "1010" в число 
// с помощью ParseInt (указана система счисления 2) и выводит результат.