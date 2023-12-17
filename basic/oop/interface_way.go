package oop

import "fmt"
type person struct{
	name string
	age int
}

func (p person) Add(i int) person{
	p.age = p.age + i
	return p
}
func interface_fn() {

	

	var p person
	p = p.Add(10)
	fmt.Println("name", p.name)

	fmt.Println(p.age)

}