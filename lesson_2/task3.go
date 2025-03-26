func groupByLength(words []string) map[int][]string {
    groups := make(map[int][]string)
    for _, word := range words {
        length := len(word)
        groups[length] = append(groups[length], word)
    }
    return groups
}