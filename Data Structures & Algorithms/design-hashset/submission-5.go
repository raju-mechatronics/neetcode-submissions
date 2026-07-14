type MyHashSet struct {
    Arr []int
}

func Constructor() MyHashSet {
    return MyHashSet{
        Arr: make([]int, 0, 100),
    }
}

func (this *MyHashSet) Add(key int) {
    for i, val := range this.Arr {
        if val == key {
            return
        }
        if key < val {
            this.Arr = append(this.Arr[:i], append([]int{key}, this.Arr[i:]...)...)
            return
        }
    }
    this.Arr = append(this.Arr, key)
}

func BinarySearch(arr []int, key int) (bool, int) {
    if len(arr) == 0 {
        return false, 0
    }
    l := 0
    r := len(arr)-1
    i := 0
    for l < r {
        i = l+(r-l)/2
        if arr[i] == key {
            return true, i
        }
        if arr[i] > key {
            r = i-1
        } else {
            l = i + 1
        }
    }
    if arr[l]==key {
        return true, l
    }
    return false, i
}

func (this *MyHashSet) Remove(key int) {
    found, i := BinarySearch(this.Arr, key)
    // fmt.Println("found", found, key, i, this.Arr)
    if found {
         this.Arr = append(this.Arr[:i], this.Arr[i+1:]...) 
    }
      
}

func (this *MyHashSet) Contains(key int) bool {
    found, _ := BinarySearch(this.Arr, key)
    return found;
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 