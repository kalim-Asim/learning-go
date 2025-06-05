package main

import ("fmt")

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob",20})
  /*  Not Works !! 
		fmt.Println(person{
			name: "A",
			age: 1
		})
	*/
	fmt.Println(person{name: "Alice",age: 19})

	s := person{name: "Sean",age: 50}

	fmt.Println(s.name) // sean
	sp := &s
	fmt.Println(sp.age) // 50
	sp.age++;
	fmt.Println(sp.age) // 51 ?
}
