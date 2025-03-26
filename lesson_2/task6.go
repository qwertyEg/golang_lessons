numbers := []int{1, 2, 2, 3, 4, 4, 5}
duplicates := make(map[int]bool)
seen := make(map[int]bool)

for _, num := range numbers {
    if /* проверьте, было ли число уже */ {
        duplicates[num] = true
    }
    // Добавьте число в seen
}

fmt.Println(duplicates) // Должно быть: map[2:true 4:true]