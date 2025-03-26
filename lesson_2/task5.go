func invertMap(original map[string]int) (map[int]string, error) {
    inverted := make(map[int]string)
    
    for key, value := range original {
        if _, exists := inverted[value]; exists {
            return nil, fmt.Errorf("дублирующееся значение: %d", value)
        }
        inverted[value] = key
    }
    
    return inverted, nil
}

// еще так можно

original := map[string]int{"a": 1, "b": 2, "c": 3}
inverted := make(map[int]string)

for key, value := range original {
    // Поменяйте ключи и значения местами
}

fmt.Println(inverted) // Должно быть: map[1:a 2:b 3:c]