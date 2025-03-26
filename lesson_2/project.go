package main

import "fmt"

const (
    CmdAdd = iota
    CmdSub
    CmdMul
    CmdDiv
    CmdPush
    CmdPop
    CmdPrint
    CmdSave
    CmdLoad
)

func main() {
    program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush, 
            -13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul, 
            CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7, 
            CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
            CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub, 
            CmdPush, 1000, CmdDiv, CmdPrint}

    stack := make([]int, 0, 100)
    registers := make(map[rune]int)
    index := 0  // счетчик

    for index < len(program) { //цикл пока у нас индекс не вышел за значения слайса program
        cmd := program[index]
        index++

        switch cmd {
        case CmdAdd:
            if len(stack) < 2 {
                panic("недостаточно элементов в стеке") //панику конечно необяз выкидывать в твоей проге
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack[len(stack)-1] = a + b

        case CmdSub:
            if len(stack) < 2 {
                panic("недостаточно элементов в стеке")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack[len(stack)-1] = a - b

        case CmdMul:
            if len(stack) < 2 {
                panic("недостаточно элементов в стеке")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack[len(stack)-1] = a * b

        case CmdDiv:
            if len(stack) < 2 {
                panic("недостаточно элементов в стеке")
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack[len(stack)-1] = a / b

        case CmdPush:
            if index >= len(program) {
                panic("нет значения для Push")
            }
            stack = append(stack, program[index])
            index++

        case CmdPop:
            if len(stack) == 0 {
                panic("Стек пуст")
            }
            stack = stack[:len(stack)-1]

        case CmdPrint:
            if len(stack) == 0 {
                panic("Стек пуст")
            }
            fmt.Println(stack[len(stack)-1])

        case CmdSave:
            if index >= len(program) {
                panic("Нет имени регистра")
            }
            if len(stack) == 0 {
                panic("Стек пуст")
            }
            reg := rune(program[index])
			index++
            registers[reg] = stack[len(stack)-1]

        case CmdLoad:
            if index >= len(program) {
                panic("Нет имени регистра")
            }
            reg := rune(program[index])
        	index++
            val, ok := registers[reg]
            if !ok {
                panic("Регистр не найден")
            }
            stack = append(stack, val)

        default:
            // Пропускаем неизвестные команды
        }
    }
}