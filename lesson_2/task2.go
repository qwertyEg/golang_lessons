func findIntersection(a, b []int) []int {
    m := make(map[int]bool)
    var result []int
    
    for _, num := range a {
        m[num] = true
    }
    
    for _, num := range b {
        if m[num] {
            result = append(result, num)
        }
    }
    return result
}