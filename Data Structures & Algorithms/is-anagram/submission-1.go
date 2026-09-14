func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    } 

    mem := make(map[rune]int)

    for _, r := range s {
        if _, ok:= mem[r]; ok {
            mem[r]++
        } else {
            mem[r] = 1
        }
    }

    for _, r:= range t {
        if _, ok:= mem[r]; ok {
            mem[r]--
            if mem[r] <= 0 {
                delete(mem, r)
            }
        } else {
            return false
        }
    }

    if len(mem) == 0 {
        return true
    }
    return false
}
