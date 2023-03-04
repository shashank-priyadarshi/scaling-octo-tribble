package main

import "fmt"

type Person struct {
	name      string
	age       int
	netIncome int
	debt      int
	country   string
	continent string
}

type Employee struct {
	Person
	salary int
}

func (p Person) NetWorth() int {
	return p.netIncome - p.debt
}

type Citizen interface {
	IsAsian() bool
	IsIndian() bool
}

type WealthManager interface {
	NetWorth() int
	IsWealthy() bool
}

type SuperRich struct {
	Person
	WealthManager
}

func (p Person) IsAsian() bool {
	return p.continent == "Asia"
}

func (p Person) IsIndian() bool {
	return p.country == "India"
}

func (p Person) IsWealthy() bool {
	return p.NetWorth() > 1
}

func (s SuperRich) IsSuperRich() {
	fmt.Printf("%v am super rich with a networth of %v", s.Person.name, s.Person.NetWorth())
}

func main() {
	fmt.Println("Hello World!")
	p := Person{name: "John", age: 30, netIncome: 50000, debt: 10000, country: "India", continent: "Asia"}
	var isSuperRich SuperRich = SuperRich{Person: p, WealthManager: p}
	isSuperRich.IsSuperRich()
}
