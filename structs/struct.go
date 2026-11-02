package main

import "fmt"

type Sport struct {
	name     string
	position string
}

type Person struct {
	name     string //a field starting with a small letter in private and cannot be accessed by other packages
	age      uint
	Height   string  // a field starting with a capital letter is public and is accessible by other packages
	sports   []Sport //creating an embedded struct
	favSport Sport
	//f    func(string) string
}

// creating a method that can be reused for every instance of a struct
func (p Person) getUser() string {
	return p.name
}

func main() {
	//a struct is a type collection of fields
	p1 := Person{age: 81, name: "Donald", favSport: Sport{"Football", "Striker"}}
	user := p1.getUser()
	fmt.Println(p1.favSport, user)
}
