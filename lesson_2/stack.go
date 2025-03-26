package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходная строка
	input := "кошка собака хомяк попугай"
	
	// Разбиваем строку на слова
	words := strings.Fields(input)
	
	// Создаем стек (используем слайс)
	stack := make([]string, 0)
	
	// Заполняем стек (Push) - кладем слова по одному
	for _, word := range words {
		stack = append(stack, word)
		fmt.Printf("Положили в стек: %s\n", word)
	}
	
	fmt.Println("\nТеперь достаем слова из стека:")
	
	// Достаем слова из стека (Pop)
	reversed := make([]string, 0)
	for len(stack) > 0 {
		// Берем последний элемент
		lastIndex := len(stack) - 1
		word := stack[lastIndex]
		
		// Удаляем его из стека
		stack = stack[:lastIndex]
		
		fmt.Printf("Достали из стека: %s\n", word)
		reversed = append(reversed, word)
	}
	
	// Собираем обратную строку
	result := strings.Join(reversed, " ")
	fmt.Println("\nРезультат:", result)
}