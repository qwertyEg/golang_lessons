func firstUniqueChar(text string) rune {
    freq := make(map[rune]int)
    
    for _, char := range text {
        freq[char]++
    }
    
    for _, char := range text {
        if freq[char] == 1 {
            return char
        }
    }
    return 0
}