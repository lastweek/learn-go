package main

import (
	"fmt"
)

type Person struct {
	age int
	sex int
}

// So by default, struct is pass by value.
// This function will not actually update the p struct
func personUpdateAge(p Person) {
	p.age = 99
}

func (p *Person) personUpdateAgeMethod() {
	p.age = 88
}

// obviously, just like C
// If we pass []int, we are already passing the pointer
func changeArrayPassByValue(a []int) {
	a[0] = 66
}
func changeArrayPassByPtr(a *[]int) {
	(*a)[1] = 77
}

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%T\n", a)

	fmt.Println(a)
	changeArrayPassByValue(a)
	changeArrayPassByPtr(&a)
	fmt.Println(a)

	// now test struct related pointer ops
	p := Person{age: 20, sex: 1}
	fmt.Println(p)

	personUpdateAge(p)
	fmt.Println(p)

	p.personUpdateAgeMethod()
	fmt.Println(p)
}
