package main

import "fmt"

type hashmap struct {
	data map[int]string
}

func NewHashmap() *hashmap {
	return &hashmap{
		data: make(map[int]string),
	}
}
func main() {

	hash := hashmap{map[int]string{}}
	hash.Set("Hello", "World")
	fmt.Println(hash)
	hash.Delete("Hello")
	fmt.Println(hash)
	fmt.Println(hash.Get("Hello"))
}

func hashstr2(key string) int {
	hash := 0
	for _, r := range key {
		hash = hash*31 + int(r)
	}
	return hash
}

func (h *hashmap) Set(key, val string) {
	hashKey := hashstr2(key)
	h.data[hashKey] = val

}

func (h *hashmap) Get(key string) (value string, ok bool) {
	hash := hashstr2(key)
	val, ok := h.data[hash]
	return val, ok
}

func (h *hashmap) Delete(key string) {
	hash := hashstr2(key)
	delete(h.data, hash)
}
