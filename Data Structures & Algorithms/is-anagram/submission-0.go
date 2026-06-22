func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    container := make(map[rune]int)
    for _, r := range s {
        if _, ok := container[r]; ok {
            container[r] = container[r]+1
        } else {
            container[r] = 1
        }
    }
    for _, u := range t {
        if _, ok := container[u]; ok {
            container[u] = container[u]-1
            if container[u] == 0 {
                delete(container, u)
            }
        } else {
            return false
        }
    }
    if len(container) > 0 {
        return false
    }
    return true
}
