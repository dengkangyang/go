package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"alice", age:30})
	fmt.Println(person{name:"Fred"})
	fmt.Println(person{name:"ann", age:40})
	
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	sp := &s
	fmt.Println(sp.age)
	sp.age = 54
	fmt.Println(sp.age)
}
