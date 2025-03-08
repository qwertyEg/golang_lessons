package main

import (
	"fmt"
	"strconv"
)

func main() {
    str := "FF"
    num, _ := strconv.ParseInt(str, 16, 64)
    fmt.Println(num)
}


//Программа преобразует шестнадцатеричную строку "FF" в число с помощью 
// ParseInt (указана система счисления 16) и выводит результат.