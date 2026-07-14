type MyHashSet struct {
    Arr []int
}

func Constructor() MyHashSet {
    return MyHashSet{
        Arr: make([]int, 0, 100),
    }
}

func (this *MyHashSet) Add(key int) {
    for _, val := range this.Arr {
        if val == key {
            return
        }
    }
    this.Arr = append(this.Arr, key)
}

func (this *MyHashSet) Remove(key int) {
    for i, val := range this.Arr {
        if val == key {
            this.Arr = append(this.Arr[:i], this.Arr[i+1:]...)
            return
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    for _, val := range this.Arr {
        if val == key {
            return true
        }
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 