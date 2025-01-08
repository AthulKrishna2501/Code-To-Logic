package main

import (
	"fmt"
)

type Entry struct {
	key   string
	value string
	next  *Entry
}
type HashTable struct {
	bucket []*Entry
	size   int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		bucket: make([]*Entry, size),
		size:   size,
	}
}

func (ht *HashTable) Hash(key string) int {
	HashValue := 0
	for _, str := range key {
		HashValue += int(str)
	}
	return HashValue % len(ht.bucket)

}

func (ht *HashTable) Insert(key, value string) {
	index := ht.Hash(key)
	newEntry := &Entry{key: key, value: value}
	if ht.bucket[index] == nil {
		ht.bucket[index] = newEntry
		return
	}

	current := ht.bucket[index]
	for current.next != nil {
		if current.key == key {
			current.value = value
		}
		current = current.next
	}
	current.next = newEntry
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.Hash(key)
	current := ht.bucket[index]
	var prev *Entry
	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.bucket[index] = current.next
			} else {
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

func (ht *HashTable) Search(key string) (string, bool) {
	index := ht.Hash(key)
	current := ht.bucket[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return "", false
}

func (ht *HashTable) Print() {
	for _, char := range ht.bucket {
		current := char
		for current != nil {
			fmt.Println(current.key, current.value)
			current = current.next
		}
	}
}
func main() {
	ht := NewHashTable(10)
	ht.Insert("Name", "Athul")
	ht.Insert("Age", "16")
	fmt.Println(ht.Search("Name"))
	ht.Print()

}
