package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
    result := make(map[string]int)
    
    // Добавляем все элементы из первого словаря
    for key, value := range map1 {
        result[key] = value
    }
    
    // Добавляем все элементы из второго словаря, перезаписывая при совпадении ключей
    for key, value := range map2 {
        result[key] = value
    }
    
    return result
}

func main() {
    map1 := map[string]int{"a": 1, "b": 2}
    map2 := map[string]int{"b": 3, "c": 4}
    fmt.Println(mergeMaps(map1, map2))
}