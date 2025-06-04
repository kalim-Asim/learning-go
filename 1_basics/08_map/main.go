package main

import ("fmt")

func main() {
	mp := make(map[string]int) // key, value
	mp["a"] = 1
	mp["b"] = 2
	
	// map[key1:value1 key2: value2]

	// _, val := mp["a"]
	val := mp["a"]
	fmt.Println(val)

	// delete(mp, "a")
	clear(mp)
	fmt.Println(mp)
}
