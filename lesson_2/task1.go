func countChars(text string) map[rune]int {
    freq := make(map[rune]int)
    for _, char := range text {
        freq[char]++
    }
    return freq
}