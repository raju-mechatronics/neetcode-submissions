

type MyHashMap struct {
	table []*list.List
}

func Constructor() MyHashMap {
	table := make([]*list.List, 10000)
    return MyHashMap{
		table : table,
	}
}

func (this *MyHashMap) Put(key int, value int) {
    hashedKey := key%10000

	if this.table[hashedKey] == nil {
		this.table[hashedKey] = list.New()
		this.table[hashedKey].PushBack([2]int{key, value})
		return
	} 


	l := this.table[hashedKey]

	for e := l.Front(); e != nil ; e = e.Next() {
		v, _ := e.Value.([2]int)
		if v[0] == key {
			e.Value = [2]int{key, value}
			return
		}
	}

	this.table[hashedKey].PushBack([2]int{key, value})	
}

func (this *MyHashMap) Get(key int) int {
    hashedKey := key%10000

	l := this.table[hashedKey]

	if l == nil {
		return -1
	}

	row := l.Front()

	for row != nil {
		v, _ := row.Value.([2]int)
		if v[0] == key {
			return v[1]
		}
		row = row.Next()
	}
	return -1

}

func (this *MyHashMap) Remove(key int) {
    hashedKey := key%10000

	l := this.table[hashedKey]

	if l == nil {
		return
	}

	row := l.Front()

	for row != nil {
		v, _ := row.Value.([2]int)
		if v[0] == key {
			l.Remove(row)
		}
		row = row.Next()
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */