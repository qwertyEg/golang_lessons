people := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 25}
ageGroups := make(map[int][]string)

for name, age := range people {
    // Заполните ageGroups: ключ - возраст, значение - имена
}

fmt.Println(ageGroups) // Пример: map[25:[Alice Charlie] 30:[Bob]]